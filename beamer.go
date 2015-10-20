package main

import (
	"github.com/denkhaus/beamer/common"
	"github.com/denkhaus/tcgl/applog"
	"github.com/lann/builder"
	"gopkg.in/pipe.v2"
)

type assembleFunc func() (pipe.Pipe, error)

var Beamer = builder.Register(beamerBuilder{}, common.BeamerCtx{}).(beamerBuilder)

type beamerBuilder builder.Builder

//############################################################################
func (p beamerBuilder) Sequence(ctxs ...common.TaskCtx) beamerBuilder {
	b := p
	for _, c := range ctxs {
		b = builder.Append(b, "Tasks", c).(beamerBuilder)
	}

	return b
}

//############################################################################
func (p beamerBuilder) Set(key string, value interface{}) beamerBuilder {
	m, ok := builder.Get(p, "KVStore")

	if ok {
		store := m.(map[string]interface{})
		store[key] = value
		return builder.Set(p, "KVStore", store).(beamerBuilder)
	}

	store := make(map[string]interface{})
	store[key] = value
	return builder.Set(p, "KVStore", store).(beamerBuilder)

}

//############################################################################
func (p beamerBuilder) On(on string) beamerBuilder {
	return builder.Set(p, "On", on).(beamerBuilder)
}

//############################################################################
func (p beamerBuilder) Build() common.BeamerCtx {
	return builder.GetStruct(p).(common.BeamerCtx)
}

//############################################################################
func (p beamerBuilder) Run() error {
	ctx := builder.GetStruct(p).(common.BeamerCtx)

	pipes := make([]pipe.Pipe, 0)
	for _, t := range ctx.Tasks {
		pip, err := t.Assemble(ctx)
		if err != nil {
			return err
		}
		pipes = append(pipes, pip)
	}

	script := pipe.Script(pipes...)
	output, err := pipe.CombinedOutput(script)
	if err != nil {
		return err
	}

	applog.Infof("%s", output)
	return nil
}
