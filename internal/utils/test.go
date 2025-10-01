package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestCheckResult(t *testing.T, name string, result, expected interface{}) {
	if result != expected {
		t.Errorf("%v: expected '%v', got '%v'", name, expected, result)
	}
}

func TestFail(t *testing.T, testName string, expected interface{}, result interface{}) {
	t.Errorf("%v: expected %v, got %v", testName, expected, result)
}

func TestGetRandomInt(min, max int) string {
	rand, _ := faker.RandomInt(min, max, 1)

	return fmt.Sprint(rand[0])
}

func TestSetEnv(env map[string]string) {
	for key, value := range env {
		os.Setenv(key, value)
	}
}

func TestUnsetEnv(env map[string]string) {
	for key, _ := range env {
		os.Unsetenv(key)
	}
}
