package godevsuite

import (
	"log/slog"
	"os"
	"path"
)

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MkdirNotExists(dir string, perm os.FileMode) error {

	dirExists, err := Exists(dir)
	if err != nil {
		return err
	}

	if !dirExists {
		if err := os.Mkdir(dir, perm); err != nil {
			return err
		}
	}

	return nil

}

const DEFAULT_FILE_PERM = os.FileMode(0664)
const DEFAULT_LOG_PERM = os.FileMode(0644)
const DEFAULT_DIR_PERM = os.FileMode(0775)

// Make sure to close the file handler when using this function ie. defer f.Close()
func SetupLogger(logPath string) (*slog.Logger, *os.File, error) {

	logDir, logFile := path.Split(logPath)

	if logDir == "" {
		logDir = "logs"
	}

	err := MkdirNotExists(logDir, DEFAULT_DIR_PERM)
	if err != nil {
		slog.Error("Log directory not created")
		slog.Error(err.Error())
		return nil, nil, err
	}

	if logFile == "" {
		var err error
		logFile, err = os.Executable()
		if err != nil {
			slog.Error("Executable name not found")
			slog.Error(err.Error())
			return nil, nil, err
		}
		_, logFile = path.Split(logFile)
		ext := path.Ext(logFile)
		logFile = logFile[0 : len(logFile)-len(ext)]
		logFile += ".log"
	}

	logPath = path.Join(logDir, logFile)

	f, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE, DEFAULT_LOG_PERM)
	if err != nil {
		slog.Error("Log file not created")
		slog.Error(err.Error())
		return nil, nil, err
	}

	logger := slog.New(slog.NewTextHandler(f, nil))
	return logger, f, nil
}
