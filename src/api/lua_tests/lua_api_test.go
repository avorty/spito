package api_tests

import (
	"testing"
	"os"
	"github.com/nasz-elektryk/spito-rules/checker"
)

func TestLuaApi(t *testing.T) {

	scripts := []string {
		"sysinfo_test.lua",
	}

	for _, script := range scripts {
		file, err := os.ReadFile(script)
		if err != nil {
			t.Fatal(err)
		}
		
		doesRulePass, err := checker.CheckRuleScript(string(file))
		if err != nil {
			t.Fatal(err)
		}

		if !doesRulePass {
			t.Fatalf("Rule %v did not pass!", file)
		}
	}
}
