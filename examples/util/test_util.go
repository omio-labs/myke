// Shared testing utilities for CLI table driven tests

package util

import (
	"github.com/goeuro/myke/cmd"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
	"strings"
	"os"
	"bytes"
	"io"
)

type TestTable struct {
	Desc     string
	Args     string
	Expected string
}

func RunCliTests(t *testing.T, dir string, tests []TestTable) {
	captureChdir(dir, func() {
		for _, tt := range tests {
			runTest(t, tt)
		}
	})
}

func runTest(t *testing.T, tt TestTable) {
	actual, err := captureStdout(func() error {
		args := strings.Split(strings.TrimSpace("myke " + tt.Args), " ")
		return cmd.NewApp().Run(args)
	})

	// TODO: Add error verification
	if assert.Regexp(t, tt.Expected, actual) {
		t.Logf("myke(%s): passed", tt.Desc)
	} else {
		t.Errorf("myke(%s): failed %s", tt.Desc, err)
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
