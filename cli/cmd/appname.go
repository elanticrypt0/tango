package cmd

import "strings"

func parseAppNameAndVersion(appname, version string) string {
	appname = strings.TrimSpace(strings.ReplaceAll(appname, " ", "-"))
	version = strings.ReplaceAll(version, ".", "-")
	return appname + "_v" + version
}
