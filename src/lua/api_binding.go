package lua

import (
	"github.com/nasz-elektryk/spito-rules/api"
	"github.com/yuin/gopher-lua"
	"reflect"
)

// Every api needs to be attached here in oder to be available:
func attachApi(L *lua.LState) {
	var t = reflect.TypeOf
	
	setGlobalConstructor(L, "Package", t(api.Package{}))
	setGlobalFunction(L, "GetCurrentDistro", api.GetCurrentDistro)
	setGlobalFunction(L, "GetCurrentDistro", api.GetCurrentDistro)
}
