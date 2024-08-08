# Linters

The same linting rules for all your teams!

## Installation

No dependency required.

Download the correct binary for your platform in the [release page](https://github.com/BlusparkTeam/linters/releases).


Some linters use Docker to run, so you need to have Docker installed on your machine.

## Usage

Initialize the configuration file:

```console
linter init
```

Edit the `.linter.yaml` file to fit your needs.

Then lint your files:

```console
linter lint
```

Note: you can initialize a pre-commit hook with:

```console
linter install
```

## Rules

Get the list of rules with:

```bash
linter rules
```


## License

AST Metrics is open-source software [licensed under the MIT license](LICENSE)

## Sponsors

![Consoneo logo](./docs/consoneo_logo.jpeg)

The digital SaaS platform for financing and managing energy renovation aid
