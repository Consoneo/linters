package rules

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
