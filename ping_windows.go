// +build windows

package main

import (
	"net"
	"os"
	"os/exec"
)

func callPingCommand(addr net.IP) error {
	// Windows automatically stops at 4 pings
	cmd := exec.Command("ping", addr.String())
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
