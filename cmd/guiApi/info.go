package guiApi

import (
	"fmt"
	"github.com/avorty/spito/pkg/shared"
	"github.com/godbus/dbus/v5"
)

type InfoApi struct {
	BusObject dbus.BusObject
}

func (i InfoApi) Log(args ...any) {
	_ = sendToDbusMethod(i.BusObject, "log", args...)
}

func (i InfoApi) Debug(args ...any) {
	_ = sendToDbusMethod(i.BusObject, "debug", args...)
}

func (i InfoApi) Error(args ...any) {
	_ = sendToDbusMethod(i.BusObject, "error", args...)
}

func (i InfoApi) Warn(args ...any) {
	_ = sendToDbusMethod(i.BusObject, "warn", args...)
}

func (i InfoApi) Important(args ...any) {
	_ = sendToDbusMethod(i.BusObject, "important", args...)
}

// Most of the time we ignore potential error because it is not really important
// and our app can work even if error is thrown
func sendToDbusMethod(busObject dbus.BusObject, logType string, values ...any) error {
	call := busObject.Call(shared.DBusMethodName("Info"), 0, logType, fmt.Sprint(values...))
	return call.Err
}
