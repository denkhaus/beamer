package tasks

import (
	"github.com/lann/builder"
)

type copyCtx struct {
	Source string
	Dest   string
}

var Copy = builder.Register(copyBuilder{}, copyCtx{}).(copyBuilder)

type copyBuilder builder.Builder

func (b copyBuilder) From(src string) copyBuilder {
	return builder.Set(b, "Source", src).(copyBuilder)
}

func (b copyBuilder) To(dest string) copyBuilder {
	return builder.Set(b, "Dest", dest).(copyBuilder)
}

func (b copyBuilder) build() copyCtx {
	return builder.GetStruct(b).(copyCtx)
}
