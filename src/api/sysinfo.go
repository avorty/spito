package api

import (
	"github.com/zcalusic/sysinfo"
	"github.com/shirou/gopsutil/v3/process"
	"strings"
	"os/exec"
)

/* TYPES */

type Distro struct {
	Name    string
	Version string
}

type InitSystem string

func (is InitSystem) String() string {
	switch is {
	case SYSTEMD:
		return "systemd"
	case RUNIT:
		return "runit"
	}
	return ""
}

func isOpenRC() bool {
	const rcCommand = "rc-status"
	cmd := exec.Command(rcCommand)
	_, err := cmd.Output()

	if err != nil {
		return false
	}

	return true
}

/* CONSTANTS */

const (
	SYSTEMD InitSystem = "systemd"
	RUNIT InitSystem = "runit"
	OPENRC InitSystem = "openrc"
	SYSV InitSystem = "sysv"
	UNKNOWN InitSystem = ""
)

/* API FUNCTIONS */

func GetDistro() Distro {
	var systemInfo sysinfo.SysInfo
	systemInfo.GetSysInfo()

	return Distro{systemInfo.OS.Name, systemInfo.OS.Release}
}

func GetInitSystem() (InitSystem, error) {
	initSystemProcess, err := process.NewProcess(1)
	if err != nil {
		return "", err
	}

	processName, err := initSystemProcess.Name()
	if err != nil {
		return "", err
	}

	if strings.Contains(processName, SYSTEMD.String()) {
		return SYSTEMD, nil
	}
	if strings.Contains(processName, RUNIT.String()) {
		return RUNIT, nil
	}
	if strings.Contains(processName, "init") {
		if isOpenRC() {
			return OPENRC, nil
		}
		return SYSV, nil
	}

	return UNKNOWN, nil
}
