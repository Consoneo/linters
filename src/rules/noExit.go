package rules

import "github.com/Consoneo/linters/src/config"

type NoExit struct {
}

func (o *NoExit) Execute(config config.Config) (string, error) {
	command := "grep -lRic --include=*.php 'exit;' " + config.Path + " ||true"
	return ExecuteCommandAndExpectNoResultToBeCorrect(command)
}

func (o *NoExit) Name() string {
	return "Check for exit() in code"
}

func (o *NoExit) Slug() string {
	return "no-exit"
}
