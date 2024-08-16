package rules

import (
	"github.com/Consoneo/linters/src/config"
)

type Psr12 struct {
}

func (o *Psr12) Execute(config config.Config) (string, error) {
	command := "docker run --rm -v " + config.Path + ":/code ghcr.io/php-cs-fixer/php-cs-fixer:${FIXER_VERSION:-3-php" + config.Version + "} check --rules=@PSR12 ."
	return ExecuteCommandAndExpectNoResultToBeCorrect(command)
}

func (o *Psr12) Name() string {
	return "Check for PSR12 compliance"
}

func (o *Psr12) Slug() string {
	return "psr12"
}

func (o *Psr12) CanFix() bool {
	return true
}

func (o *Psr12) Fix(config config.Config) (string, error) {
	command := "docker run --rm -v " + config.Path + ":/code ghcr.io/php-cs-fixer/php-cs-fixer:${FIXER_VERSION:-3-php" + config.Version + "} fix --rules=@PSR12 ."
	return ExecuteCommandAndExpectNoResultToBeCorrect(command)
}

