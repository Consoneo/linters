package rules

import (
	"os"
	"os/exec"

	"github.com/Consoneo/linters/src/config"
	log "github.com/sirupsen/logrus"
)

type AstMetrics struct {
}

func (o *AstMetrics) Execute(c config.Config) (string, error) {

	// Downloading binary if needed
	binaryPath := os.TempDir() + string(os.PathSeparator) + "ast-metrics"

	fileExists, err := os.Stat(binaryPath)
	if err != nil || fileExists == nil {
		log.Debug("AstMetrics binary not found, downloading it")
		// to /tmp/ast-metrics
		command := "curl -s https://raw.githubusercontent.com/Halleck45/ast-metrics/main/scripts/download.sh|bash && chmod +x ast-metrics && mv ast-metrics " + binaryPath
		log.Debug("Executing command: ", command)
		cmd := exec.Command("sh", "-c", command)
		out, err := cmd.CombinedOutput()
		if err != nil {
			return string(out), err
		}
	}

	command := binaryPath + " analyze --non-interactive --report-html=build/astmetrics " + c.Path
	log.Debug("Executing command: ", command)

	r, err := ExecuteCommandAndExpectNoResultToBeCorrect(command)

	return r, err
}

func (o *AstMetrics) Name() string {
	return "Runs AstMetrics static analysis"
}

func (o *AstMetrics) Slug() string {
	return "ast-metrics"
}
