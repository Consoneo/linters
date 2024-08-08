package rules

import "github.com/Consoneo/linters/src/config"

type Psr1 struct {
}

func (o *Psr1) Execute(config config.Config) (string, error) {
	command := "docker run --rm -v " + config.Path + ":/data cytopia/phpcs --standard=PSR1 ."
	return ExecuteCommandAndExpectNoResultToBeCorrect(command)
}

func (o *Psr1) Name() string {
	return "PSR1"
}

func (o *Psr1) Slug() string {
	return "psr1"
}
