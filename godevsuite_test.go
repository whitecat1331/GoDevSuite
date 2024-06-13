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
	logger, f, err := SetupSLogger("")
	if err != nil {
		t.Log("SLogger Failed")
		t.Fatal(err.Error())
	}
	defer f.Close()
	logger.Info("SLogger Setup Successfully")
}
func TestSetupSLoggerWithFile(t *testing.T) {
	logger, f, err := SetupSLogger("test.log")
	if err != nil {
		t.Log("SLogger Failed")
		t.Fatal(err.Error())
	}
	defer f.Close()
	logger.Info("SLogger Setup Successfully")
}

func TestSetupSLoggerWithPath(t *testing.T) {
	logger, f, err := SetupSLogger(path.Join("test_logs", "test.log"))
	if err != nil {
		t.Log("SLogger Failed")
		t.Fatal(err.Error())
	}
	defer f.Close()
	logger.Info("SLogger Setup Successfully")
}
func TestSetupSLoggerAsLog(t *testing.T) {
	logger, f, err := SetupSLogger(path.Join("test_logs", "test.log"))
	if err != nil {
		t.Log("SLogger Failed")
		t.Fatal(err.Error())
	}
	defer f.Close()
	logger.Info("SLogger Setup Successfully")
}
