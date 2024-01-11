package dbutils

import "os"

func GetPhotoDir() string {
	wd, _ := os.Getwd()
	return wd + "\\photos"
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
