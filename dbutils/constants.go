package dbutils

import "os"

var PhotoFolder = "photos"

func GetPhotoDir() string {
	wd, _ := os.Getwd()
	return wd + "\\" + PhotoFolder
}

var Offset = 0
var Limit = -1

func SetDefaults(offset **int, limit **int) {
	if *offset == nil {
		*offset = &Offset
	}
	if *limit == nil {
		*limit = &Limit
	}
}
