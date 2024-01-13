package fileutils

import "strings"

func Generatephotofilename(tigername string, lastseen string) string {
	ret := strings.ReplaceAll(tigername+"_"+lastseen, " ", "_")
	return strings.ReplaceAll(ret, ":", "_")
}
