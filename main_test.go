package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/bitrise-io/go-utils/command"
	"github.com/bitrise-io/go-utils/env"
	"github.com/bitrise-io/go-xcode/utility"
	"github.com/bitrise-io/go-xcode/xcarchive"
)

func TestConfig_validate(t *testing.T) {
	type fields struct {
		ArchivePath                     string
		ExportMethod                    string
		UploadBitcode                   bool
		CompileBitcode                  bool
		TeamID                          string
		CustomExportOptionsPlistContent string
		UseLegacyExport                 bool
		DeployDir                       string
		VerboseLog                      bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    fields
		wantErr bool
	}{
		{
			name: "TeamID contains whitespace",
			fields: fields{
				TeamID: "  ",
			},
			want: fields{
				TeamID: "",
			},
			wantErr: false,
		},
		{
			name: "ExportOptionsPlistContent contains whitespace",
			fields: fields{
				CustomExportOptionsPlistContent: "  ",
			},
			want: fields{
				CustomExportOptionsPlistContent: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configs := &Config{
				ArchivePath:               tt.fields.ArchivePath,
				DistributionMethod:        tt.fields.ExportMethod,
				UploadBitcode:             tt.fields.UploadBitcode,
				CompileBitcode:            tt.fields.CompileBitcode,
				TeamID:                    tt.fields.TeamID,
				ExportOptionsPlistContent: tt.fields.CustomExportOptionsPlistContent,
				DeployDir:                 tt.fields.DeployDir,
				VerboseLog:                tt.fields.VerboseLog,
			}
			wantConfigs := &Config{
				ArchivePath:               tt.want.ArchivePath,
				DistributionMethod:        tt.want.ExportMethod,
				UploadBitcode:             tt.want.UploadBitcode,
				CompileBitcode:            tt.want.CompileBitcode,
				TeamID:                    tt.want.TeamID,
				ExportOptionsPlistContent: tt.want.CustomExportOptionsPlistContent,
				DeployDir:                 tt.want.DeployDir,
				VerboseLog:                tt.want.VerboseLog,
			}
			if err := configs.validate(); (err != nil) != tt.wantErr {
				t.Errorf("Config.validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(configs, wantConfigs) {
				t.Errorf("Config.validate() configs = %+v, wantConfig = %+v", configs, wantConfigs)
			}
		})
	}
}

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

	print("%s", result)
}

func TestConfig_generateExportOptions_plistValidField(t *testing.T) {
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
