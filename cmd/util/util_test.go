package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestDirExists_ok(t *testing.T) {
	testcases := []struct {
		dirExists bool
		expected  bool
	}{
		{
			dirExists: true,
			expected:  true,
		},
		{
			dirExists: false,
			expected:  false,
		},
	}

	for _, tc := range testcases {
		var dir string

		if tc.dirExists {
			d, err := ioutil.TempDir("", "TestDirExistsOK")
			if err != nil {
				t.Fatal(err)
			}
			defer os.RemoveAll(d)

			dir = d
		} else {
			abs, err := filepath.Abs(".")
			if err != nil {
				t.Fatal(err)
			}

			dir = filepath.Join(abs, "doesnotexist")
		}

		got, err := DirExists(dir)
		if err != nil {
			t.Errorf("got error: %s", err)
		}

		if got != tc.expected {
			t.Errorf("expected: %t, got: %t", tc.expected, got)
		}
	}
}

func TestDirExists_error(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "TestDirExistsError")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	_, gotErr := DirExists(tmpfile.Name())
	if gotErr == nil {
		t.Error("error should raise")
	}

	expected := fmt.Sprintf("%q is a file", tmpfile.Name())
	if gotErr.Error() != expected {
		t.Errorf("error expected: %q, got: %q", expected, gotErr.Error())
	}
}
