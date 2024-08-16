package rules

import "github.com/Consoneo/linters/src/config"

type PhpStan struct {
}

func (o *PhpStan) Execute(config config.Config) (string, error) {
	command := "docker run --rm -v "+ config.Path +":/app ghcr.io/phpstan/phpstan analyse ."
	return ExecuteCommandAndExpectNoResultToBeCorrect(command)
}

func (o *PhpStan) Name() string {
	return "Run PHPStan analysis"
}

func (o *PhpStan) Slug() string {
	return "phpstan"
}
