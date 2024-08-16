package rules

import "github.com/Consoneo/linters/src/config"

type Symfony struct {
}

func (o *Symfony) Execute(config config.Config) (string, error) {
	command := "docker run --rm -v " + config.Path + ":/code ghcr.io/php-cs-fixer/php-cs-fixer:${FIXER_VERSION:-3-php" + config.Version + "} check --rules=@PSR12 ."
	return ExecuteCommandAndExpectNoResultToBeCorrect(command)
}

func (o *Symfony) Name() string {
	return "Check for PHPCS @Symfony rules compliance"
}

func (o *Symfony) Slug() string {
	return "symfony"
}
