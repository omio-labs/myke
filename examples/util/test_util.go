// Shared testing utilities for CLI table driven tests

package util

import (
	"bytes"
	"fmt"
	"github.com/goeuro/myke/cmd"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestTable represents a table-driven test
type TestTable struct {
	Arg string
	Out string
	Err bool
}

// RunCliTests runs myke CLI with the given table tests
func RunCliTests(t *testing.T, dir string, tests []TestTable) {
	os.Setenv("COLUMNS", "999")
	captureChdir(dir, func() {
		for _, tt := range tests {
			runTest(t, tt)
		}
	})
}

func runTest(t *testing.T, tt TestTable) {
	actual, err := captureStdout(func() error {
		args := strings.Split(tt.Arg, " ")
		return cmd.Exec(args)
	})

	if err != nil {
		// add error message to output for validating error messages
		actual = fmt.Sprintf("%v %v", actual, err)
	}

	expectedOut := strings.Replace(tt.Out, "$PS$", "/", -1)
	expectedOut = strings.Replace(expectedOut, "$PLS$", ":", -1)

	if tt.Err == (err != nil) && assert.Regexp(t, expectedOut, actual) {
		t.Logf("myke(%s): passed", tt.Arg)
	} else {
		t.Errorf("myke(%s): failed %s", tt.Arg, err)
	}
}

func captureChdir(dir string, f func()) {
	olddir, _ := os.Getwd()
	os.Chdir(filepath.Join(olddir, dir))
	f()
	os.Chdir(olddir)
}

func captureStdout(f func() error) (string, error) {
	oldStdout := os.Stdout
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()

	os.Stdout = w
	os.Stderr = w
	err := f()
	os.Stdout = oldStdout
	os.Stderr = oldStderr

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String(), err
}
