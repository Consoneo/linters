package rules

import "github.com/Consoneo/linters/src/config"

type CustomCommand struct {
	Command string
}

func (o *CustomCommand) Execute() (string, error) {
	return ExecuteCommandAndExpectNoResultToBeCorrect(o.Command)
}

func (o *CustomCommand) Name() string {
	return o.Command
}

func (o *CustomCommand) Slug() string {
	return "custom"
}

func (o *CustomCommand) CanFix() bool {
	return false
}

func (o *CustomCommand) Fix(config config.Config) (string, error) {
	return "", nil
}
