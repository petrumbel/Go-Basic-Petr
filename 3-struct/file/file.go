package file

import (
	"os"
	"strings"
)

func IsJSON(filename string) bool {
	return strings.HasSuffix(strings.ToLower(filename), ".json")
}

func EnsureFileExists(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = file.WriteString("[]")
		return err
	}
	return nil
}
