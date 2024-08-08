package engine

import (
	"fmt"
	"os"
	"sync"

	"github.com/Consoneo/linters/src/config"
	"github.com/Consoneo/linters/src/rules"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/chelnak/ysmrr"
	"gopkg.in/yaml.v2"
)

type Analyse struct{}

func (o *Analyse) Lint() {
	lintFile := ".linters.yaml"

	// check if the file exists
	if _, err := os.Stat(lintFile); os.IsNotExist(err) {
		fmt.Println("File .linters.yaml does not exist")
		os.Exit(1)
	}

	// Load the .linters.yaml file
	file, err := os.Open(lintFile)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(1)
	}
	defer file.Close()

	// Unmarshal the yaml file
	var linters map[string]interface{}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&linters)
	if err != nil {
		fmt.Println("Error decoding yaml: ", err)
		os.Exit(1)
	}

	// Ensure docker is installed
	_, err = rules.CheckIfDockerIsInstalled()
	if err != nil {
		fmt.Println("Docker is not installed")
		os.Exit(1)
	}

	// Start spinner
	sm := ysmrr.NewSpinnerManager()
	sm.Start()
	var wg sync.WaitGroup

	languages := []string{"php", "javascript", "css", "html"}
	for _, language := range languages {

		// check if language is configured
		if _, ok := linters["lints"].(map[interface{}]interface{})[language]; !ok {
			continue
		}

		yamlConfig := linters["lints"].(map[interface{}]interface{})[language]
		directories := yamlConfig.(map[interface{}]interface{})["src"].([]interface{})

		// if key version exists
		version := "latest"
		if _, ok := yamlConfig.(map[interface{}]interface{})["version"]; ok {
			version = yamlConfig.(map[interface{}]interface{})["version"].(string)
		}

		pattern := "*"
		switch language {
		case "php":
			pattern = "*.php"
		case "javascript":
			pattern = "*.js"
			if version == "latest" {
				version = "22"
			}
		case "css":
			pattern = "*.css"
		case "html":
			pattern = "*.html"
		}

		config := config.Config{
			Version: version,
			Pattern: pattern,
		}

		for _, dir := range directories {

			// if not rules, skip
			if _, ok := yamlConfig.(map[interface{}]interface{})["rules"]; !ok {
				continue
			}

			// get rules
			rulesToApply := yamlConfig.(map[interface{}]interface{})["rules"].([]interface{})
			ch := make(chan int, len(rulesToApply))

			availableRules := rules.Rules()
			for _, availableRule := range availableRules {

				for _, rule := range rulesToApply {

					// if not applyable (same slug, pass)
					if availableRule.Slug() != rule {
						continue
					}

					wg.Add(1)
					go func(rule interface{}) {

						spinner := sm.AddSpinner("(" + language + ") " + availableRule.Name())

						config.Path = dir.(string)
						_, err := availableRule.Execute(config)
						if err != nil {
							spinner.Error()
						} else {
							spinner.Complete()
						}

						wg.Done()
						ch <- 1
					}(rule)
				}
			}
		}
	}

	wg.Wait()
	sm.Stop()
}

func (o *Analyse) InitConfig() {
	// create a .linters.yaml file with basic configuration
	lintFile := ".linters.yaml"
	if _, err := os.Stat(lintFile); !os.IsNotExist(err) {
		fmt.Println(".linters.yaml already exists")
		os.Exit(1)
	}

	file, err := os.Create(lintFile)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		os.Exit(1)
	}

	defer file.Close()

	template := `
	lints:
  php:
    version: "8.4"
    src:
      - src
      - src2
    rules:
      - no-syntax-error
      - no-dump
      - no-exit
      - psr1
      - psr2
      
  javascript:
    version: "8.1"
    src:
      - ./front/
    rules:
      - eslint
`

	_, err = file.WriteString(template)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		os.Exit(1)
	}

	fmt.Println("File .linters.yaml created")
}

func (o *Analyse) Install() {

	// Ensure current dir is a git repository
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		fmt.Println("Not a git repository")
		os.Exit(1)
	}

	// create a pre-commit hook
	hookFile := ".git/hooks/pre-commit"
	if _, err := os.Stat(hookFile); !os.IsNotExist(err) {
		fmt.Println("pre-commit hook already exists")
		os.Exit(1)
	}

	file, err := os.Create(hookFile)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		os.Exit(1)
	}

	defer file.Close()

	template := `#!/bin/sh
	linters lint
	`

	_, err = file.WriteString(template)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		os.Exit(1)
	}

	err = os.Chmod(hookFile, 0755)
	if err != nil {
		fmt.Println("Error changing permissions: ", err)
		os.Exit(1)
	}

	fmt.Println("pre-commit hook created")
}

func (o *Analyse) ListRules() {

	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		Headers("RULE", "DESCRIPTION")


	availableRules := rules.Rules()
	for _, availableRule := range availableRules {
		t.Row(availableRule.Slug(), availableRule.Name())
	}

	fmt.Println(t.Render())
}
