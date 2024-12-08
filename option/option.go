package option

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

type RunMode int

const (
	Development RunMode = iota
	Staging
	Production
)

func (r RunMode) String() string {
	switch r {
	case Development:
		return "development"
	case Staging:
		return "staging"
	case Production:
		return "production"
	default:
		return "development"
	}
}

type Option struct {
	RunMode RunMode
	File    struct {
		BasePath string
		Name     map[RunMode]string
	}
}

var (
	once     sync.Once
	instance *Option
)

func NewOption() *Option {
	once.Do(func() {
		instance = &Option{
			File: struct {
				BasePath string
				Name     map[RunMode]string
			}{
				Name: make(map[RunMode]string),
			},
		}
	})
	return instance
}

func (opt *Option) SetFileName(mode RunMode, filename string) *Option {
	opt.File.Name[mode] = filename
	return opt
}

func (opt *Option) SetRunMode(mode RunMode) *Option {
	opt.RunMode = mode
	return opt
}

func (opt *Option) SetBasePath(path string) *Option {
	opt.File.BasePath = path
	return opt
}

func (opt *Option) LoadCurrentRunMode(key string) *Option {
	mode := os.Getenv(key)
	fmt.Println(mode)
	switch {
	case strings.EqualFold(mode, Development.String()):
		opt.RunMode = Development
	case strings.EqualFold(mode, Staging.String()):
		opt.RunMode = Staging
	case strings.EqualFold(mode, Production.String()):
		opt.RunMode = Production
	default:
		opt.RunMode = Development
	}
	return opt
}
