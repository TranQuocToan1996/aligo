package support

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"github.com/essentialkaos/ek/v12/fmtutil"
	"github.com/essentialkaos/ek/v12/system"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// showOSInfo shows verbose information about system
func showOSInfo() {
	systemInfo, err := system.GetSystemInfo()

	if err != nil {
		return
	}

	fmtutil.Separator(false, "SYSTEM INFO")

	printInfo(7, "Name", systemInfo.OS)
	printInfo(7, "Version", systemInfo.Version)
	printInfo(7, "Arch", systemInfo.Arch)
	printInfo(7, "Kernel", systemInfo.Kernel)
}

// collectEnvInfo collects info about environment
func collectEnvInfo() Pkgs {
	return nil
}

// showEnvInfo shows info about environment
func showEnvInfo(pkgs Pkgs) {
	return
}
