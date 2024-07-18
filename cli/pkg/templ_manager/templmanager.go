package templmanager

import "fmt"

func GetRenderArgs(rootpath string) string {
	return fmt.Sprintf("generate -path %s/api/app/views ", rootpath)
}
