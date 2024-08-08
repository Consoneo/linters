package rules

import "github.com/Consoneo/linters/src/config"

type Symfony struct {
}

func (o *Symfony) Execute(config config.Config) (string, error) {
	command := "docker run --rm -v " + config.Path + ":/data cytopia/phpcs --standard=Symfony ."
	return ExecuteCommandAndExpectNoResultToBeCorrect(command)
}

func (o *Symfony) Name() string {
	return "Symfony coding style"
}

func (o *Symfony) Slug() string {
	return "symfony"
}
