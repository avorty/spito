package tests

import (
	"github.com/avorty/spito/pkg/vrct"
	"github.com/avorty/spito/pkg/vrct/vrctFs"
	"os"
	"path/filepath"
	"testing"
)

type ConfigsSetup struct {
	configs         []Config
	resultPath      string
	destinationPath string
	configType      vrctFs.FileType
}

type Config struct {
	path        string
	optionsPath string
	isOptional  bool
}

func TestConfigsMatrix(t *testing.T) {
	ruleVrct, err := vrct.NewRuleVRCT()
	if err != nil {
		t.Fatal("Failed to Create VRCT instance")
	}
	fsVrct := &ruleVrct.Fs

	tmpPath, err := os.MkdirTemp("/tmp", "spito-test-")
	if err != nil {
		t.Fatal("Failed to create temporary test directory\n", err)
	}

	configs := []ConfigsSetup{
		{
			configs: []Config{
				{
					path:       "json/eslint-default.json",
					isOptional: true,
				},
				{
					path:        "json/eslint-esprima.json",
					optionsPath: "json/esprima-options.json",
					isOptional:  false,
				},
			},
			resultPath:      "json/eslint-merged.json",
			destinationPath: tmpPath + "/new_dir/eslint.json",
			configType:      vrctFs.JsonConfig,
		},
		{
			configs: []Config{
				{
					path:       "yaml/extrepo-default.yaml",
					isOptional: true,
				},
				{
					path:       "yaml/extrepo-full.yaml",
					isOptional: false,
				},
			},
			resultPath:      "yaml/extrepo-full.yaml",
			destinationPath: tmpPath + "/new_dir/extrepo.yaml",
			configType:      vrctFs.YamlConfig,
		},
		{
			configs: []Config{
				{
					path:       "toml/hugo-default.toml",
					isOptional: false,
				},
				{
					path:       "toml/hugo-customized.toml",
					isOptional: true,
				},
			},
			resultPath:      "toml/hugo-merged.toml",
			destinationPath: tmpPath + "/new_dir/hugo.toml",
			configType:      vrctFs.TomlConfig,
		},
	}

	for _, config := range configs {
		testConfigs(t, fsVrct, config)
	}

	// cleanup
	_ = os.RemoveAll(tmpPath)
}

func testConfigs(t *testing.T, vrct *vrctFs.VRCTFs, setup ConfigsSetup) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to obtain working directory: '%s'", wd)
	}
	wd = filepath.Join(wd, "config_data")

	for _, config := range setup.configs {
		workingPath := filepath.Join(wd, config.path)
		configTestData, err := os.ReadFile(workingPath)
		if err != nil {
			t.Fatalf("Failed to open test data '%s': %s", workingPath, err)
		}

		workingOptionsPath := filepath.Join(wd, config.optionsPath)

		var options []byte
		if config.optionsPath != "" {
			options, err = os.ReadFile(workingOptionsPath)
			if err != nil {
				t.Fatalf("Failed to open result data '%s': %s", workingOptionsPath, err)
			}
		}

		err = vrct.CreateConfig(setup.destinationPath, configTestData, options, config.isOptional, setup.configType)
		if err != nil {
			t.Fatal("Failed trying to override file "+setup.destinationPath+"\n", err)
		}
	}
	workingResPath := filepath.Join(wd, setup.resultPath)
	desiredRawResult, err := os.ReadFile(workingResPath)
	if err != nil {
		t.Fatalf("Failed to open result data '%s': %s", workingResPath, err)
	}

	obtainedRawResult, err := vrct.ReadFile(setup.destinationPath)
	if err != nil {
		t.Fatalf("Failed to read file destinationPath %s: %s", setup.destinationPath, err)
	}

	// It is empty because we are not reverting anything
	_, err = vrct.Apply([]vrctFs.Rule{}, false)
	if err != nil {
		t.Fatal("Failed to apply VRCT\n", err)
	}

	obtainedRealRawResult, err := os.ReadFile(setup.destinationPath)
	if err != nil {
		t.Fatalf("Failed to read from real fs file '%s': %s", setup.destinationPath, err)
	}

	err = vrctFs.CompareConfigs(obtainedRawResult, desiredRawResult, setup.configType)
	if err != nil {
		t.Fatalf("Failed to properly merge virtual fs file '%s': %s", setup.destinationPath, err)
	}

	err = vrctFs.CompareConfigs(obtainedRealRawResult, desiredRawResult, setup.configType)
	if err != nil {
		t.Fatalf("Failed to properly save real fs file '%s': %s", setup.destinationPath, err)
	}

}
