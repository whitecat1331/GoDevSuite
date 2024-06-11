package godevsuite

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"strings"
	// "testing"
)

type TestValues struct {
	Domain  string
	Domains []string
	Output  string
}

func NewTestingValues() *TestValues {
	domains := strings.Split(os.Getenv("DOMAINS"), ",")
	return &TestValues{
		Domain:  domains[0],
		Domains: domains,
		Output:  os.Getenv("OUTPUT"),
	}

}

var TV = NewTestingValues()

// func TestFunctionName(t *testing.T) {
//
// }
