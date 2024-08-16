package rules

func Rules() []Rule {
	return []Rule{
		// php
		&NoDump{},
		&NoSyntaxError{},
		&NoExit{},
		&Psr12{},
		&Psr1{},
		&Psr2{},
		&Symfony{},
		&PhpStan{},
		&PhpCS{},
		// javascript
		&EsLint{},
		// agnostic
		&AstMetrics{},
	}
}
