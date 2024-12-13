package file

import "path/filepath"

func Ext(path string) string {
	return filepath.Ext(path)
}
