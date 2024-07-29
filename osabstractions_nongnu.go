//go:build windows
// +build windows

package main

import "io/fs"

func init() {
	panic("I dont work on Windows")
}

// CopyFS copies a folder and contents to another location
// Waiting on os.CopyFS to be implemented.
// Not hacked around in Windows as not sure on syntax
func CopyFS(destDir string, fsys fs.FS) error {
	return nil
}
