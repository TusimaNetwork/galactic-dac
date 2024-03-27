package near

import (
	"os/exec"
)

func DoCommand(command string, args ...string) (string, error) {
	//cmd := exec.Command("node", "-e", "console.log('Hello from Node.js')")
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
