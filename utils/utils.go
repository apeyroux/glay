package utils

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
)

type ExecResult struct {
	Output []byte
	Err    error
}

func Run(cmd string, rchan chan ExecResult) {

	fcmd := strings.Fields(cmd)

	go func() {
		r := ExecResult{}
		r.Err = nil // TODO: obligatoire ? à tester
		xcmd := exec.Command(fcmd[0], fcmd[1:len(fcmd)]...)
		output, err := xcmd.Output()
		if err != nil {
			r.Err = err
		}
		r.Output = output
		rchan <- r
	}()
}

func PidIsAlive(pid int) (state bool) {
	p, err := os.FindProcess(pid)

	if err != nil {
		return false
	}

	err = p.Signal(syscall.Signal(0))

	if err != nil && err.Error() == "no such process" {
		return false
	}
	return true
}
