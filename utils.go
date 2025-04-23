package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"github.com/bitrise-io/go-xcode/v2/exportoptionsgenerator"
)

// ParseExportProduct ...
func ParseExportProduct(product string) (exportoptionsgenerator.ExportProduct, error) {
	switch product {
	case "app":
		return exportoptionsgenerator.ExportProductApp, nil
	case "app-clip":
		return exportoptionsgenerator.ExportProductAppClip, nil
	default:
		return "", fmt.Errorf("unkown method (%s)", product)
	}
}

func findIDEDistrubutionLogsPath(output string) (string, error) {
	pattern := `IDEDistribution: -\[IDEDistributionLogging _createLoggingBundleAtPath:\]: Created bundle at path '(?P<log_path>.*)'`
	re := regexp.MustCompile(pattern)

	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		if match := re.FindStringSubmatch(line); len(match) == 2 {
			return match[1], nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", nil
}
