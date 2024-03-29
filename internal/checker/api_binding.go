package checker

import (
	"reflect"

	"github.com/avorty/spito/pkg/api"
	"github.com/avorty/spito/pkg/shared"
	"github.com/avorty/spito/pkg/vrct/vrctFs"
	"github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

// Every cmdApi needs to be attached here to be available:
func attachApi(importLoopData *shared.ImportLoopData, ruleConf *shared.RuleConfigLayout, L *lua.LState) {
	apiNamespace := newLuaNamespace()

	apiNamespace.AddField("pkg", getPackageNamespace(importLoopData, L))
	apiNamespace.AddField("sys", getSysInfoNamespace(L))
	apiNamespace.AddField("daemon", getDaemonApiNamespace(importLoopData, L))
	apiNamespace.AddField("fs", getFsNamespace(importLoopData, L))
	apiNamespace.AddField("info", getInfoNamespace(importLoopData, L))
	apiNamespace.AddField("git", getGitNamespace(importLoopData, L))

	if ruleConf.Unsafe {
		apiNamespace.AddField("sh", getShNamespace(L))
	}

	apiNamespace.setGlobal(L, "api")
}

func getPackageNamespace(importLoopData *shared.ImportLoopData, L *lua.LState) lua.LValue {
	pkgNamespace := newLuaNamespace()
	pkgNamespace.AddFn("get", api.GetPackage)
	pkgNamespace.AddFn("install", func(packagesToInstall ...string) error {
		for _, packageToCheck := range packagesToInstall {
			err := importLoopData.PackageTracker.AddPackage(packageToCheck)
			if err != nil {
				return err
			}
		}
		return api.InstallPackages(packagesToInstall...)
	})
	pkgNamespace.AddFn("remove", func(packagesToRemove ...string) error {
		for _, packageToCheck := range packagesToRemove {
			err := importLoopData.PackageTracker.RemovePackage(packageToCheck)
			if err != nil {
				return err
			}
		}
		return api.RemovePackages(packagesToRemove...)
	})

	return pkgNamespace.createTable(L)
}

func getSysInfoNamespace(L *lua.LState) lua.LValue {
	sysInfoNamespace := newLuaNamespace()

	sysInfoNamespace.AddFn("sleep", api.Sleep)
	sysInfoNamespace.AddFn("getDistro", api.GetDistro)
	sysInfoNamespace.AddFn("getInitSystem", api.GetInitSystem)
	sysInfoNamespace.AddFn("getRandomLetters", api.GetRandomLetters)
	sysInfoNamespace.AddFn("getEnv", api.GetEnv)

	return sysInfoNamespace.createTable(L)
}

func getDaemonApiNamespace(importLoopData *shared.ImportLoopData, L *lua.LState) lua.LValue {
	daemonNamespace := newLuaNamespace()

	daemonApi := api.DaemonApi{ImportLoopData: importLoopData}

	daemonNamespace.AddFn("start", daemonApi.StartDaemon)
	daemonNamespace.AddFn("stop", daemonApi.StopDaemon)
	daemonNamespace.AddFn("restart", daemonApi.RestartDaemon)
	daemonNamespace.AddFn("enable", daemonApi.EnableDaemon)
	daemonNamespace.AddFn("disable", daemonApi.DisableDaemon)

	daemonNamespace.AddFn("get", api.GetDaemon)

	return daemonNamespace.createTable(L)
}

func getFsNamespace(importLoop *shared.ImportLoopData, L *lua.LState) lua.LValue {
	fsNamespace := newLuaNamespace()

	apiFs := api.FsApi{FsVRCT: &importLoop.VRCT.Fs}

	fsNamespace.AddFn("pathExists", apiFs.PathExists)
	fsNamespace.AddFn("fileExists", apiFs.FileExists)
	fsNamespace.AddFn("readFile", apiFs.ReadFile)
	fsNamespace.AddFn("readDir", apiFs.ReadDir)
	fsNamespace.AddFn("fileContains", api.FileContains)
	fsNamespace.AddFn("removeComments", api.RemoveComments)
	fsNamespace.AddFn("find", api.Find)
	fsNamespace.AddFn("findAll", api.FindAll)
	fsNamespace.AddFn("getProperLines", api.GetProperLines)
	fsNamespace.AddFn("createFile", apiFs.CreateFile)
	fsNamespace.AddFn("createConfig", apiFs.CreateConfig)
	fsNamespace.AddFn("updateConfig", apiFs.UpdateConfig)
	fsNamespace.AddFn("compareConfigs", apiFs.CompareConfigs)
	fsNamespace.AddFn("copy", apiFs.Copy)
	fsNamespace.AddFn("apply", apiFs.Apply)
	fsNamespace.AddField("config", getConfigEnums(L))

	return fsNamespace.createTable(L)
}

func getConfigEnums(L *lua.LState) lua.LValue {
	infoNamespace := newLuaNamespace()

	infoNamespace.AddField("json", lua.LNumber(vrctFs.JsonConfig))
	infoNamespace.AddField("yaml", lua.LNumber(vrctFs.YamlConfig))
	infoNamespace.AddField("toml", lua.LNumber(vrctFs.TomlConfig))

	return infoNamespace.createTable(L)
}

func getInfoNamespace(importLoopData *shared.ImportLoopData, L *lua.LState) lua.LValue {
	infoApi := importLoopData.InfoApi
	infoNamespace := newLuaNamespace()

	infoNamespace.AddFn("log", infoApi.Log)
	infoNamespace.AddFn("debug", infoApi.Debug)
	infoNamespace.AddFn("warn", infoApi.Warn)
	infoNamespace.AddFn("error", infoApi.Error)
	infoNamespace.AddFn("important", infoApi.Important)

	return infoNamespace.createTable(L)
}

func getGitNamespace(importLoopData *shared.ImportLoopData, L *lua.LState) lua.LValue {
	gitNamespace := newLuaNamespace()

	gitApi := api.GitApi{FsVrct: &importLoopData.VRCT.Fs}

	gitNamespace.AddFn("clone", gitApi.GitClone)

	return gitNamespace.createTable(L)
}

func getShNamespace(L *lua.LState) lua.LValue {
	shellNamespace := newLuaNamespace()

	shellNamespace.AddFn("command", api.ShellCommand)
	shellNamespace.AddFn("exec", api.Exec)

	return shellNamespace.createTable(L)
}

type LuaNamespace struct {
	functions map[string]interface{}
	fields    map[string]lua.LValue
}

func newLuaNamespace() LuaNamespace {
	return LuaNamespace{
		functions: make(map[string]interface{}),
		fields:    make(map[string]lua.LValue),
	}
}

func (ln LuaNamespace) AddFn(name string, fn interface{}) {
	ln.functions[name] = mapFunctionErrorReturnToString(fn)
}

func (ln LuaNamespace) AddField(name string, field lua.LValue) {
	ln.fields[name] = field
}

func (ln LuaNamespace) setGlobal(L *lua.LState, name string) {
	namespaceTable := ln.createTable(L)
	L.SetGlobal(name, namespaceTable)
}

func (ln LuaNamespace) createTable(L *lua.LState) *lua.LTable {
	namespaceTable := L.NewTable()

	for fnName, fn := range ln.functions {
		L.SetField(namespaceTable, fnName, luar.New(L, fn))
	}
	for fieldName, field := range ln.fields {
		L.SetField(namespaceTable, fieldName, field)
	}

	return namespaceTable
}

func isTypeError(t reflect.Type) bool {
	errorType := reflect.TypeOf((*error)(nil)).Elem()
	return t.Implements(errorType)
}

func mapFunctionErrorReturnToString(fn any) any {
	fnType := reflect.TypeOf(fn)
	reflectLValueType := reflect.TypeOf((*lua.LValue)(nil)).Elem()

	if fnType.Kind() != reflect.Func {
		// It throws panic, because if I would avoid it, fnType.NumIn() would panic which would be harder to debug
		panic("fn argument in `mapFunctionErrorReturnToString` must be a function")
	}

	var inTypes = make([]reflect.Type, fnType.NumIn())
	var outTypes = make([]reflect.Type, fnType.NumOut())

	for i := 0; i < fnType.NumIn(); i++ {
		inTypes[i] = fnType.In(i)
	}

	for i := 0; i < fnType.NumOut(); i++ {
		outType := fnType.Out(i)

		// If returns error, map it to string (lua.LString)
		if isTypeError(outType) {
			outTypes[i] = reflectLValueType
		} else {
			outTypes[i] = outType
		}
	}

	// Create new function definition
	newFnType := reflect.FuncOf(inTypes, outTypes, fnType.IsVariadic())

	return reflect.MakeFunc(newFnType, func(args []reflect.Value) []reflect.Value {
		var fnResults []reflect.Value

		// IDK why but .Call doesn't automatically detect if it is variadic
		if newFnType.IsVariadic() {
			fnResults = reflect.ValueOf(fn).CallSlice(args)
		} else {
			fnResults = reflect.ValueOf(fn).Call(args)
		}

		var newResults = make([]reflect.Value, len(fnResults))

		for i, result := range fnResults {
			// If not error - skip
			if !isTypeError(result.Type()) {
				newResults[i] = result
				continue
			}
			var newErr lua.LValue

			// If error is not nil, map it to string
			if result.Interface() == nil {
				newErr = lua.LNil
			} else {
				err := result.Interface().(error).Error()
				newErr = lua.LString(err)
			}

			newResults[i] = reflect.ValueOf(newErr)
		}

		return newResults
	}).Interface()
}
