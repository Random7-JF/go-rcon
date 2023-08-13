package validator

import (
	"fmt"
	"strings"
)

func isBlank(field string) bool {
	if len(field) != 0 {
		return true
	} else {
		return false
	}
}

func hasValue(cmd string, value string) (bool, error) {
	if isBlank(value) {
		return false, fmt.Errorf("%s command requires value, got no value: %s", cmd, value)
	}
	return true, nil
}

func hasOption(cmd string, option string, validOptions []string) (bool, error) {
	if isBlank(option) {
		return false, fmt.Errorf("%s command requires and option, got no option", cmd)
	}
	for _, opt := range validOptions {
		if strings.EqualFold(option, opt) {
			return true, nil
		}
	}
	return false, fmt.Errorf("%s command has %s as valid options but recieved %s as option", cmd, validOptions, option)
}
