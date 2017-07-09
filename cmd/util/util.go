package util

import (
	"os"

	"github.com/pkg/errors"
)

// DirExists return the given directory exists or not
func DirExists(dir string) (bool, error) {
	stat, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, errors.Wrapf(err, "failed to check whether %q exists or not", dir)
	}

	if !stat.IsDir() {
		return false, errors.Errorf("%q is a file", dir)
	}

	return true, nil
}
