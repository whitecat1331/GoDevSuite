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

func TestSetupLoggerEmpty(t *testing.T) {
	logger, f, err := SetupLogger("")
	if err != nil {
		t.Log("Logger Failed")
		t.Fatal(err.Error())
	}
	defer f.Close()
	logger.Info("Logger Setup Successfully")
}

func TestSetupLoggerWithFile(t *testing.T) {
	logger, f, err := SetupLogger("test.log")
	if err != nil {
		t.Log("Logger Failed")
		t.Fatal(err.Error())
	}
	defer f.Close()
	logger.Info("Logger Setup Successfully")
}

func TestSetupLoggerWithPath(t *testing.T) {
	logger, f, err := SetupLogger(path.Join("test_logs", "test.log"))
	if err != nil {
		t.Log("Logger Failed")
		t.Fatal(err.Error())
	}
	defer f.Close()
	logger.Info("Logger Setup Successfully")
}
