package main

import (
	"strings"
	"testing"

	"github.com/bitrise-io/go-utils/v2/command"
	"github.com/bitrise-io/go-utils/v2/env"
	"github.com/bitrise-io/go-xcode/utility"
	"github.com/bitrise-io/go-xcode/xcarchive"
	"github.com/stretchr/testify/assert"
)

func TestConfig_generateExportOptions_plist(t *testing.T) {
	// Given
	envRepository := env.NewRepository()
	commandFactory := command.NewFactory(envRepository)
	xcodebuildVersion, _ := utility.GetXcodeVersion(commandFactory)
	archive, _ := xcarchive.NewIosArchive("configs.ArchivePath")

	// When
	result, _ := generateExportOptionsPlist("app", "development", "my team id", false, false, xcodebuildVersion.MajorVersion, archive, false)

	// Then
	if len(result) == 0 {
		t.Errorf("plist is empty")
	}

	assert.NotEqual(t, 0, len(result))
}

func TestConfig_generateExportOptions_plist_validField(t *testing.T) {
	// Given
	envRepository := env.NewRepository()
	commandFactory := command.NewFactory(envRepository)
	xcodebuildVersion, _ := utility.GetXcodeVersion(commandFactory)
	archive, _ := xcarchive.NewIosArchive("configs.ArchivePath")

	// When
	result, err := generateExportOptionsPlist("app", "development", "my team id", false, false, xcodebuildVersion.MajorVersion, archive, true)

	// Then
	assert.Nil(t, err)
	assert.Contains(t, result, "compileBitcode")
	assert.Contains(t, result, "method")
	assert.Contains(t, result, "development")
}

func TestConfig_generateExportOptions_plist_updateVersionAndBuildSetToFalse(t *testing.T) {
	// Given
	envRepository := env.NewRepository()
	commandFactory := command.NewFactory(envRepository)
	xcodebuildVersion, _ := utility.GetXcodeVersion(commandFactory)
	archive, _ := xcarchive.NewIosArchive("configs.ArchivePath")

	// When
	result, err := generateExportOptionsPlist("app", "app-store", "my team id", false, false, xcodebuildVersion.MajorVersion, archive, false)

	// Then
	assert.Nil(t, err)

	if xcodebuildVersion.MajorVersion > 12 && strings.Contains(result, "manageAppVersionAndBuildNumber") == false {
		t.Errorf("plist does not contain manage app version and build number value for method field")
	}
}
