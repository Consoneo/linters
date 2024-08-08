package rules

import "github.com/Consoneo/linters/src/config"

type Psr12 struct {
}

func (o *Psr12) Execute(config config.Config) (string, error) {
	command := "docker run --rm -v " + config.Path + ":/data cytopia/phpcs --standard=PSR12 ."
	return ExecuteCommandAndExpectNoResultToBeCorrect(command)
}

func (o *Psr12) Name() string {
	return "PSR12"
}

func (o *Psr12) Slug() string {
	return "psr12"
}
