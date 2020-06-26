package modules

import (
	"fmt"
	"strings"
)

func pass(key string) (string, error) {
	result := Shell("pass " + key)()
	if strings.HasPrefix("Error: ", result) {
		return "", fmt.Errorf(result)
	}
	return result, nil
}
