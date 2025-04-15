package main

/*
func newXcodeVersionReader(t *testing.T, major int64) xcodeversion.Reader {
	mockXcodeVersionReader := mocks.NewXcodeVersionReader(t)
	mockXcodeVersionReader.On("GetVersion").Return(xcodeversion.Version{
		Major: major,
	})

	return mockXcodeVersionReader
}


func TestConfig_generateExportOptions_plist(t *testing.T) {
	// Given
	xcodebuildVersion, _ := utility.GetXcodeVersion()
	xcodeVersionReaeder := newXcodeVersionReader(t, 15)
	archive, _ := xcarchive.NewIosArchive("configs.ArchivePath")

	g := exportoptionsgenerator.NewWithIosArchive(archive, exportoptionsgenerator.ExportProductApp, xcodeVersionReaeder, log.NewLogger())

	// When
	result, _ := g.GenerateApplicationExportOptions("app", "development", "my team id", false, false, , false)

	// Then
	if len(result) == 0 {
		t.Errorf("plist is empty")
	}

	assert.NotEqual(t, 0, len(result))
}

func TestConfig_generateExportOptions_plist_validField(t *testing.T) {
	// Given
	xcodebuildVersion, _ := utility.GetXcodeVersion()
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
	xcodebuildVersion, _ := utility.GetXcodeVersion()
	archive, _ := xcarchive.NewIosArchive("configs.ArchivePath")

	// When
	result, err := generateExportOptionsPlist("app", "app-store", "my team id", false, false, xcodebuildVersion.MajorVersion, archive, false)

	// Then
	assert.Nil(t, err)

	if xcodebuildVersion.MajorVersion > 12 && strings.Contains(result, "manageAppVersionAndBuildNumber") == false {
		t.Errorf("plist does not contain manage app version and build number value for method field")
	}
}
*/
