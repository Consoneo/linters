package rules

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func ExecuteCommandAndExpectNoResultToBeCorrect(command string) (string, error) {

	log.Debug("Running command: ", command)
	// execute the command
	cmd := exec.Command("bash", "-c", command)
	stdout, err := cmd.Output()
	if err != nil {
		log.Debug("command fails: ", command)
		return "", err
	}

	log.Debug("Command output: ", string(stdout))

	// if stdout is empty, returns OK
	if string(stdout) == "" {
		return "", nil
	}

	// if stdout is not empty, returns the output
	return string(command), nil
}

func CheckIfDockerIsInstalled() (string, error) {
	return ExecuteCommandAndExpectNoResultToBeCorrect("docker --version")
}