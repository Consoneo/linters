package rules

import "github.com/Consoneo/linters/src/config"

type Rule interface {
	Slug() string

	Name() string

	Execute(config config.Config) (string, error)
}
