package rules

import "github.com/Consoneo/linters/src/config"

type NoDump struct {
}

func (o *NoDump) Execute(config config.Config) (string, error) {
	command := "grep -lRic --include=*.php 'dump(' " + config.Path + " ||true"
	return ExecuteCommandAndExpectNoResultToBeCorrect(command)
}

func (o *NoDump) Name() string {
	return "No var_dump in code"
}

func (o *NoDump) Slug() string {
	return "no-dump"
}
