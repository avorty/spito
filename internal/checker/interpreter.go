package checker

import (
	"github.com/avorty/spito/pkg/shared"
	"github.com/yuin/gopher-lua"
)

const rulesetDirConstantName = "RULESET_DIR"

func ExecuteLuaMain(script string, importLoopData *shared.ImportLoopData, ruleConf *RuleConf, rulesetPath string) (bool, error) {
	L := lua.NewState(lua.Options{SkipOpenLibs: true})
	defer L.Close()

	// Standard libraries
	lua.OpenString(L)

	L.SetGlobal(rulesetDirConstantName, lua.LString(rulesetPath))
	attachApi(importLoopData, ruleConf, L)
	attachRuleRequiring(importLoopData, ruleConf, L)

	if err := L.DoString(script); err != nil {
		return false, err
	}

	err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("main"),
		Protect: true,
		NRet:    1,
	})
	if err != nil {
		return false, err
	}

	return bool(L.Get(-1).(lua.LBool)), nil
}

func attachRuleRequiring(importLoopData *shared.ImportLoopData, ruleConf *RuleConf, L *lua.LState) {
	L.SetGlobal("require_rule", L.NewFunction(func(state *lua.LState) int {
		ruleUrl := L.Get(1).String()
		ruleName := L.Get(2).String()

		result := _internalCheckRule(importLoopData, ruleUrl, ruleName, ruleConf)
		L.Push(lua.LBool(result))

		return 1
	}))
}
