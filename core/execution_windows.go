// +build windows

package core

import (
	"os"
	"os/exec"
	"path/filepath"
)

var comSpec string

func init() {
	comSpec = os.Getenv("ComSpec")
	if len(comSpec) > 0 {
		return
	}
	winDir := os.Getenv("windir")
	if len(winDir) > 0 {
		comSpec = filepath.Join(winDir, "system32", "cmd.exe")
	}
	comSpec = "C:\\windows\\system32\\cmd.exe"
	return
}

func executionShell() []string {
	return []string{comSpec, "/C"}
}

func (e *Execution) beforeExecuteCmd(cmd string, env map[string]string) error {
	if len(e.Task.Shell) > 0 {
		return nil
	}
	// This will cause the same output like the sh -x on unix like systems.
	proc := exec.Command(comSpec, "/C", "echo", "+", cmd)
	proc.Dir = e.Project.Cwd
	proc.Env = mapToSlice(env)
	proc.Stdout = os.Stdout
	proc.Stderr = os.Stderr
	return proc.Run()
}
