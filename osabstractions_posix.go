//go:build darwin || freebsd || linux || netbsd || openbsd || solaris || !windows

package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// CopyFS copies a folder and contents to another location
// Waiting on os.CopyFS to be implemented...
// until that time. Eww.
func CopyFS(dir string, fsys fs.FS) error {
	return fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		targ := filepath.Join(dir, filepath.FromSlash(path))
		if d.IsDir() {
			if err := os.MkdirAll(targ, 0777); err != nil {
				return err
			}
			return nil
		}
		r, err := fsys.Open(path)
		if err != nil {
			return err
		}
		defer r.Close()
		info, err := r.Stat()
		if err != nil {
			return err
		}
		w, err := os.OpenFile(targ, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666|info.Mode()&0777)
		if err != nil {
			return err
		}
		if _, err := io.Copy(w, r); err != nil {
			w.Close()
			return fmt.Errorf("copying %s: %v", path, err)
		}
		if err := w.Close(); err != nil {
			return err
		}
		return nil
	})
}
