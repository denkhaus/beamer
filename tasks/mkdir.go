package tasks

import (
	"github.com/denkhaus/tcgl/applog"
	"github.com/lann/builder"
)

type dirCtx struct {
	Directory   string
	Description string
}

var dir = builder.Register(dirBuilder{}, dirCtx{}).(dirBuilder)

type dirBuilder builder.Builder

func (b dirBuilder) Create(dir string) dirBuilder {
	return builder.Set(b, "Directory", dir).(dirBuilder)
}

func (b dirBuilder) Descr(descr string) dirBuilder {
	return builder.Set(b, "Description", descr).(dirBuilder)
}

func (b dirBuilder) PrintDescription() {
	val, ok := builder.Get(b, "Description")
	if !ok {
		applog.Warningf("Mkdir has no description")
		return
	}
	applog.Infof(val.(string))
}

func (b dirBuilder) build() dirCtx {
	return builder.GetStruct(b).(dirCtx)
}

func Mkdir(directory string) dirBuilder {
	return dir.Create(directory)
}
