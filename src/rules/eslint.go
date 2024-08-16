package rules

import "github.com/Consoneo/linters/src/config"

type EsLint struct {
}

func (o *EsLint) Execute(config config.Config) (string, error) {
	command := "docker run --rm --workdir=/data -v " + config.Path + ":/data node:" + config.Version + "-alpine node_modules/.bin/eslint /data"
	return ExecuteCommandAndExpectNoResultToBeCorrect(command)
}

func (o *EsLint) Name() string {
	return "Run ESLint"
}

func (o *EsLint) Slug() string {
	return "eslint"
}
