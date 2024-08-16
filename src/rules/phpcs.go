package rules

import (
	"github.com/Consoneo/linters/src/config"
)

type PhpCS struct {
}

func (o *PhpCS) Execute(config config.Config) (string, error) {
	command := "docker run --rm -v " + config.Path + ":/code ghcr.io/php-cs-fixer/php-cs-fixer:${FIXER_VERSION:-3-php" + config.Version + "} check --rules=@PSR12 ."
	return ExecuteCommandAndExpectNoResultToBeCorrect(command)
}

func (o *PhpCS) Name() string {
	return "Run PHP CS Fixer"
}

func (o *PhpCS) Slug() string {
	return "phpcs"
}
