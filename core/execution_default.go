// +build !windows

package core

func executionShell() []string {
	return []string{"sh", "-exc"}
}

func (e *Execution) beforeExecuteCmd(cmd string, env map[string]string) error {
	return nil
}
