package rules

import "github.com/Consoneo/linters/src/config"

type Psr2 struct {
}

func (o *Psr2) Execute(config config.Config) (string, error) {
	command := "docker run --rm -v " + config.Path + ":/data cytopia/phpcs --standard=PSR2 ."
	return ExecuteCommandAndExpectNoResultToBeCorrect(command)
}

func (o *Psr2) Name() string {
	return "PSR2"
}

func (o *Psr2) Slug() string {
	return "psr2"
}
