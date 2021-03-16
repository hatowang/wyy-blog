package main

import (
	"os/exec"
	"fmt"
)

func main() {
	runCmd()
}

func runCmd() {
	output, _ := exec.Command( "/bin/sh", []string{"-c", "kubectl top pods -n kube-system metrics-server-74dbd77757-2kchl"}...).CombinedOutput()
	fmt.Println(string(output))
}
