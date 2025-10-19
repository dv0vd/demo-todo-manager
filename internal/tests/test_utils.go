package testutils

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/go-faker/faker/v4"
)

func CheckResult(t *testing.T, name string, result, expected interface{}) {
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("%v: expected '%v', got '%v'", name, expected, result)
	}
}

func Fail(t *testing.T, testName string, expected interface{}, result interface{}) {
	t.Errorf("%v: expected %v, got %v", testName, expected, result)
}

func GetRandomInt(min, max int) string {
	rand, _ := faker.RandomInt(min, max, 1)

	return fmt.Sprint(rand[0])
}

func SetEnv(env map[string]string) {
	for key, value := range env {
		os.Setenv(key, value)
	}
}

func UnsetEnv(env map[string]string) {
	for key, _ := range env {
		os.Unsetenv(key)
	}
}
