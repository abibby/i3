package modules

import (
	"fmt"
	"strings"
)

func pass(key string) (string, error) {
	result := Shell("pass " + key + " | head -n 1")()
	if strings.HasPrefix("Error: ", result) {
		return "", fmt.Errorf(result)
	}
	return result, nil
}
