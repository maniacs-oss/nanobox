// +build !windows

package util

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jcelliott/lumber"
	"github.com/nanobox-io/nanobox-golang-stylish"
)

// PrivilegeExec runs a command as sudo
func PrivilegeExec(command, msg string) {
	fmt.Printf(stylish.Bullet(msg))

	//
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo %v %v", os.Args[0], command))

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// run command
	if err := cmd.Run(); err != nil {
		lumber.Fatal("[util/util_unix]", err.Error())
	}
}
