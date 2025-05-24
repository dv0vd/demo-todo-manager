package utils

import (
	"fmt"

	"github.com/go-faker/faker/v4"
)

func TestGetRandomInt(min, max int) string {
	rand, _ := faker.RandomInt(min, max, 1)

	return fmt.Sprint(rand[0])
}
