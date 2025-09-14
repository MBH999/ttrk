package version

// Version contains the version information for the application
const Version = "0.1.0"

// BuildDate contains the build date
var BuildDate string

// GitCommit contains the git commit hash
var GitCommit string

// GetVersion returns the full version information
func GetVersion() string {
	version := Version
	if GitCommit != "" {
		version += " (" + GitCommit + ")"
	}
	if BuildDate != "" {
		version += " built on " + BuildDate
	}
	return version
}