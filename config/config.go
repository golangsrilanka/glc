package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
	"path"
)

var (
	GLCDir     string
	ConfigPath string
	LogFile    string
)

func init() {
	home, err := homedir.Dir()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "fatal: cannot find home directory: %v", err)
		os.Exit(1)
	}

	GLCDir = path.Join(home, ".glc")
	ConfigPath = path.Join(GLCDir, "config.yaml")
	LogFile = path.Join(GLCDir, "debug-log")
}

// InitDebugLogFile initializes the debug-log file (create/append)
// and returns a pointer to the file.
func InitDebugLogFile() (*os.File, error) {
	f, err := os.OpenFile(LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening debug-log file: %w", err)
	}
	return f, nil
}

func InitGLCDirectory() error {
	var err error
	if _, err = os.Stat(ConfigPath); err == nil {
		return nil
	}

	if !os.IsNotExist(err) {
		return fmt.Errorf("unexpected error: %w", err)
	}

	// Config file not found; initialize new directory
	if err := os.Mkdir(GLCDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create %s directory: %w", GLCDir, err)
	}

	// Create new config file
	if _, err = os.Create(ConfigPath); err != nil {
		return fmt.Errorf("failed to create new config file %s: %w", ConfigPath, err)
	}

	return nil
}

//TODO: Set Provider not working
func SetGitProvider(selectedProvider string) {
	viper.Set("provider", selectedProvider)
}
