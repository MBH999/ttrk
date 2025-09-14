package version

import (
	"strings"
	"testing"
)

func TestGetVersion(t *testing.T) {
	// Test with default values
	version := GetVersion()
	if !strings.Contains(version, Version) {
		t.Errorf("Expected version to contain %s, got %s", Version, version)
	}

	// Test with build date
	BuildDate = "2024-01-01:12:00:00"
	version = GetVersion()
	if !strings.Contains(version, BuildDate) {
		t.Errorf("Expected version to contain build date %s, got %s", BuildDate, version)
	}

	// Test with git commit
	GitCommit = "abc123"
	version = GetVersion()
	if !strings.Contains(version, GitCommit) {
		t.Errorf("Expected version to contain git commit %s, got %s", GitCommit, version)
	}

	// Reset values
	BuildDate = ""
	GitCommit = ""
}