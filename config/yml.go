package config

import (
	"path"
	"path/filepath"
	"runtime"
)

type YAML struct {
}

func (y *YAML) PathFile() string {
	return RootDir() + "/config.yml"
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
