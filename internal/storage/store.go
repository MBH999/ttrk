package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/MBH999/ttrk/internal/config"
)

// Data represents the on-disk structure storing projects and tasks.
type Data struct {
	Projects []Project `json:"projects"`
}

// Project groups several tasks under a single name.
type Project struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Tasks []Task `json:"tasks"`
}

// Task represents a unit of work tracked inside a project.
type Task struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	TotalSeconds int64       `json:"total_seconds"`
	Entries      []TimeEntry `json:"entries"`
}

// TimeEntry records one contiguous timer session.
type TimeEntry struct {
	ID              string    `json:"id"`
	TaskID          string    `json:"task_id"`
	Start           time.Time `json:"start"`
	End             time.Time `json:"end"`
	DurationSeconds int64     `json:"duration_seconds"`
}

// Store manages persistence of tracker data on disk.
type Store struct {
	path string
}

// DefaultStore initialises a store beneath the user's configuration directory.
func DefaultStore() (*Store, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	if cfg.DataDir == "" {
		return nil, errors.New("data directory not configured")
	}

	if err := os.MkdirAll(cfg.DataDir, 0o755); err != nil {
		return nil, fmt.Errorf("create storage directory: %w", err)
	}

	return &Store{path: filepath.Join(cfg.DataDir, "data.json")}, nil
}

// NewStore builds a Store rooted at the supplied file path. The parent directory must exist.
func NewStore(path string) *Store {
	return &Store{path: path}
}

// Path exposes the destination file for tests or callers.
func (s *Store) Path() string {
	return s.path
}

// Load reads tracker data from disk, returning an empty dataset if the file does not exist.
func (s *Store) Load() (Data, error) {
	var data Data

	contents, err := os.ReadFile(s.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return data, nil
		}
		return data, fmt.Errorf("read storage file: %w", err)
	}

	if len(contents) == 0 {
		return data, nil
	}

	if err := json.Unmarshal(contents, &data); err != nil {
		return data, fmt.Errorf("decode storage file: %w", err)
	}

	return data, nil
}

// Save persists the supplied data structure to disk.
func (s *Store) Save(data Data) error {
	if s.path == "" {
		return errors.New("storage path unset")
	}

	payload, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("encode storage data: %w", err)
	}

	tmp := s.path + ".tmp"
	if err := os.WriteFile(tmp, payload, 0o600); err != nil {
		return fmt.Errorf("write temp storage file: %w", err)
	}

	if err := os.Rename(tmp, s.path); err != nil {
		return fmt.Errorf("replace storage file: %w", err)
	}

	return nil
}

// AddEntry registers a completed timer session against the task.
func (t *Task) AddEntry(start, end time.Time) time.Duration {
	if end.Before(start) {
		start, end = end, start
	}
	duration := end.Sub(start)
	seconds := duration / time.Second
	if seconds <= 0 {
		seconds = 1
	}
	duration = time.Duration(seconds) * time.Second

	entry := TimeEntry{
		ID:              NewID(),
		TaskID:          t.ID,
		Start:           start,
		End:             end,
		DurationSeconds: int64(seconds),
	}

	t.Entries = append(t.Entries, entry)
	t.TotalSeconds += int64(seconds)

	return duration
}

// TotalDuration returns the cumulative tracked time for the task.
func (t Task) TotalDuration() time.Duration {
	return time.Duration(t.TotalSeconds) * time.Second
}

// NewID produces a simple unique identifier.
func NewID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
