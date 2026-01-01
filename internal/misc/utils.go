package misc

import (
	"fmt"
	"os"
	"strings"
)

func GetWorkingFolder() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "spring-boot-app", err
	}
	return dir, nil

}

func ValidateNoSpaces(str string) error {
	if str != "" && strings.Contains(str, " ") {
		return fmt.Errorf("Name cannot contain spaces")
	}
	return nil
}
