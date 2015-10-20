package main

import "github.com/lann/builder"

type TaskCtx interface {
	PrintDescription()
}

type beamerCtx struct {
	Tasks []TaskCtx
}

var Beamer = builder.Register(beamerBuilder{}, beamerCtx{}).(beamerBuilder)

type beamerBuilder builder.Builder

func (p beamerBuilder) Sequence(ctxs ...TaskCtx) {
	for c := range ctxs {
		builder.Append(p, "Tasks", c)
	}
}

func (p beamerBuilder) Run() {
	st := builder.GetStruct(p).(beamerCtx)
	for _, t := range st.Tasks {

		t.PrintDescription()
	}
}
