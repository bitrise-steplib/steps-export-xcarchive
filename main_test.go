package main

import (
	"strings"
	"testing"

	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-io/go-utils/env"
	"github.com/bitrise-io/go-xcode/utility"
	"github.com/bitrise-io/go-xcode/xcarchive"
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
	if err != nil {
		t.Errorf("generate export options plist error")
	}

	if strings.Contains(result, "compileBitcode") == false {
		t.Errorf("plist does not contain compile bitcode field")
	}

	if strings.Contains(result, "method") == false {
		t.Errorf("plist does not contain method field")
	}

	if strings.Contains(result, "development") == false {
		t.Errorf("plist does not contain development value for method field")
	}
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
	if err != nil {
		t.Errorf("generate export options plist error")
	}

	if strings.Contains(result, "manageAppVersionAndBuildNumber") == false {
		t.Errorf("plist does not contain manage app version and build number value for method field")
	}
}
