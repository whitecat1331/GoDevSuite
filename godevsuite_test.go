package godevsuite

import (
	"path"
	"testing"
)

func TestTestingTemplate(t *testing.T) {
	t.Logf("DOMAINS: %#v", TV.Domains)
	t.Logf("DOMAIN: %#v", TV.Domain)
	t.Logf("OUTPUT: %#v", TV.Output)
}

func TestSetupSLoggerEmpty(t *testing.T) {
	logger, f := SetupSLogger("")
	err := f()
	if err != nil {
		t.Log("Failed to close file")
		t.Fatal(err.Error())
	}
	logger.Info("SLogger Setup Successfully")
}
func TestSetupSLoggerWithFile(t *testing.T) {
	logger, f := SetupSLogger("test.log")
	defer f()
	logger.Info("SLogger Setup Successfully")
}

func TestSetupSLoggerWithPath(t *testing.T) {
	logger, f := SetupSLogger(path.Join("test_logs", "test.log"))
	defer f()
	logger.Info("SLogger Setup Successfully")
}
func TestSetupSLoggerAsLog(t *testing.T) {
	logger, f := SetupSLogger(path.Join("test_logs", "test.log"))
	defer f()
	logger.Info("SLogger Setup Successfully")
}
