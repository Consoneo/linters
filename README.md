# Linters

The same linting rules for all your teams!

## Installation

No dependency required.

```console
curl -s https://raw.githubusercontent.com/BlusparkTeam/linters/main/scripts/download.sh|bash
chmod +x linters
```

Or download the correct binary for your platform in the [release page](https://github.com/BlusparkTeam/linters/releases).


Some linters use Docker to run, so you need to have Docker installed on your machine.

## Getting started

Generate a `.linters.yaml` the configuration file:

```console
linters init
```

Edit the `.linter.yaml` file to fit your needs.

For example:

```yaml
lints:
  php:
    version: "8.1"
    src:
      - src1
      - my_another_directory
    rules:
      - no-syntax-error
      - no-dump
      - no-exit
      - psr12
```

Then lint your files:

```console
linters lint
```

## Auto-fix

Some linters can fix the code for you:

```console
linters fix
```

## Integration with git


You can initialize a pre-commit (by default) hook with:

```console
linters install
```
You can specify hooks to create with options `--pre-commit`, `--pre-push`

## Rules

Get the list of rules with:

```bash
linters rules
```

Today, the following linters are available:


| Name              | Description                               |
|-------------------|-------------------------------------------|
| `no-dump`         | Check for var_dump in code                |
| `no-syntax-error` | Check for syntax errors in PHP files      |
| `no-exit`         | Check for exit() in code                  |
| `psr12`           | Check for PSR12 compliance                |
| `psr1`            | Check for PSR1 compliance                 |
| `psr2`            | Check for PSR2 compliance                 |
| `symfony`         | Check for PHPCS @Symfony rules compliance |
| `phpstan`         | Run PHPStan analysis                      |
| `phpcs`           | Run PHP CS Fixer                          |
| `eslint`          | Run ESLint                                |
| `ast-metrics`     | Runs AstMetrics static analysis           |



## License

Linters is open-source software [licensed under the MIT license](LICENSE)

## Sponsors

![Consoneo logo](./docs/consoneo_logo.jpeg)

The digital SaaS platform for financing and managing energy renovation aid
