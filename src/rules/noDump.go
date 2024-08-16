package rules

import "github.com/Consoneo/linters/src/config"

type NoDump struct {
}

func (o *NoDump) Execute(config config.Config) (string, error) {
	command := "grep -lRic --include=*.php 'dump(' " + config.Path + " ||true"
	return ExecuteCommandAndExpectNoResultToBeCorrect(command)
}

func (o *NoDump) Name() string {
	return "Check for var_dump in code"
}

func (o *NoDump) Slug() string {
	return "no-dump"
}

func (o *NoDump) CanFix() bool {
	return false
}

func (o *NoDump) Fix(config config.Config) (string, error) {
	return "", nil
}
