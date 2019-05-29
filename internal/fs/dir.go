package fs

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func MakeDirectory(path string) error {
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("skipping, path exists %s", path)
	}

	err := os.MkdirAll(path, 0755)
	if err != nil {
		return errors.Wrap(err, "MakeDirectory")
	}
	return nil
}

func MakeDirs(dirs []string) error {
	for _, folder := range dirs {
		if err := MakeDirectory(folder); err != nil {
			return err
		}
	}

	return nil
}
