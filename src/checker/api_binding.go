package checker

import (
	"github.com/nasz-elektryk/spito/api"
	"github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
	"reflect"
)

// Every api needs to be attached here in order to be available:
func attachApi(L *lua.LState) {
	var t = reflect.TypeOf

	apiNamespace := newLuaNamespace()

	setGlobalConstructor(L, "Package", t(api.Package{}))

	apiNamespace.AddField("sys", getSysInfoNamespace(L))
	apiNamespace.AddField("fs", getFsNamespace(L))

	apiNamespace.setGlobal(L, "api")
}

func getSysInfoNamespace(L *lua.LState) lua.LValue {
	sysInfoNamespace := newLuaNamespace()

	sysInfoNamespace.AddFn("GetDistro", api.GetDistro)
	sysInfoNamespace.AddFn("GetDaemon", api.GetDaemon)
	sysInfoNamespace.AddFn("GetInitSystem", api.GetInitSystem)

	return sysInfoNamespace.createTable(L)
}

func getFsNamespace(L *lua.LState) lua.LValue {
	fsNamespace := newLuaNamespace()

	fsNamespace.AddFn("PathExists", api.PathExists)
	fsNamespace.AddFn("PathExists", api.PathExists)
	fsNamespace.AddFn("FileExists", api.FileExists)
	fsNamespace.AddFn("ReadFile", api.ReadFile)
	fsNamespace.AddFn("ReadDir", api.ReadDir)
	fsNamespace.AddFn("FileContains", api.FileContains)
	fsNamespace.AddFn("RemoveComments", api.RemoveComments)
	fsNamespace.AddFn("Find", api.Find)
	fsNamespace.AddFn("FindAll", api.FindAll)
	fsNamespace.AddFn("GetProperLines", api.GetProperLines)

	return fsNamespace.createTable(L)
}

type LuaNamespace struct {
	constructors map[string]reflect.Type
	functions    map[string]interface{}
	fields       map[string]lua.LValue
}

func newLuaNamespace() LuaNamespace {
	return LuaNamespace{
		constructors: map[string]reflect.Type{},
		functions:    make(map[string]interface{}),
		fields:       make(map[string]lua.LValue),
	}
}

func (ln LuaNamespace) AddConstructor(name string, Obj reflect.Type) {
	ln.constructors[name] = Obj
}

func (ln LuaNamespace) AddFn(name string, fn interface{}) {
	ln.functions[name] = fn
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
	for constrName, constrInterface := range ln.constructors {
		constr := constructorFunction(L, constrInterface)
		L.SetField(namespaceTable, constrName, constr)
	}
	for fieldName, field := range ln.fields {
		L.SetField(namespaceTable, fieldName, field)
	}

	return namespaceTable
}

func setGlobalConstructor(L *lua.LState, name string, Obj reflect.Type) {
	L.SetGlobal(name, constructorFunction(L, Obj))

}

func constructorFunction(L *lua.LState, Obj reflect.Type) lua.LValue {
	return L.NewFunction(func(state *lua.LState) int {
		obj := reflect.New(Obj)

		state.Push(luar.New(state, obj.Interface()))
		return 1
	})
}