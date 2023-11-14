package api_tests

import (
	"github.com/nasz-elektryk/spito/checker"
	"os"
	"testing"
)

func TestLuaApi(t *testing.T) {

	scripts := []string{
		"sysinfo_test.lua",
		"fs_rules_test.lua",
		"rule_require_test.lua",
	}

	for _, scriptName := range scripts {
		file, err := os.ReadFile(scriptName)
		if err != nil {
			t.Fatal(err)
		}

		doesRulePass, err := checker.CheckRuleScript(string(file))
		if err != nil {
			t.Fatal(err)
		}

		if !doesRulePass {
			t.Fatalf("Rule %s did not pass!", scriptName)
		}
	}
}