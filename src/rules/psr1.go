package rules

import "github.com/Consoneo/linters/src/config"

type Psr1 struct {
}

func (o *Psr1) Execute(config config.Config) (string, error) {
	command := "docker run --rm -v " + config.Path + ":/code ghcr.io/php-cs-fixer/php-cs-fixer:${FIXER_VERSION:-3-php" + config.Version + "} check --rules=@PSR1 ."
	return ExecuteCommandAndExpectNoResultToBeCorrect(command)
}

func (o *Psr1) Name() string {
	return "Check for PSR1 compliance"
}

func (o *Psr1) Slug() string {
	return "psr1"
}

func (o *Psr1) CanFix() bool {
	return true
}

func (o *Psr1) Fix(config config.Config) (string, error) {
	command := "docker run --rm -v " + config.Path + ":/code ghcr.io/php-cs-fixer/php-cs-fixer:${FIXER_VERSION:-3-php" + config.Version + "} fix --rules=@PSR1 ."
	return ExecuteCommandAndExpectNoResultToBeCorrect(command)
}

