package rules

import "github.com/Consoneo/linters/src/config"

type Psr2 struct {
}

func (o *Psr2) Execute(config config.Config) (string, error) {
	command := "docker run --rm -v " + config.Path + ":/code ghcr.io/php-cs-fixer/php-cs-fixer:${FIXER_VERSION:-3-php" + config.Version + "} check --rules=@PSR2 ."
	return ExecuteCommandAndExpectNoResultToBeCorrect(command)
}

func (o *Psr2) Name() string {
	return "Check for PSR2 compliance"
}

func (o *Psr2) Slug() string {
	return "psr2"
}

func (o *Psr2) CanFix() bool {
	return true
}

func (o *Psr2) Fix(config config.Config) (string, error) {
	command := "docker run --rm -v " + config.Path + ":/code ghcr.io/php-cs-fixer/php-cs-fixer:${FIXER_VERSION:-3-php" + config.Version + "} fix --rules=@PSR2 ."
	return ExecuteCommandAndExpectNoResultToBeCorrect(command)
}

