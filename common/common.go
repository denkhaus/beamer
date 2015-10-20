package common

import (
	"bufio"
	"bytes"
	"text/template"

	"gopkg.in/pipe.v2"
)

type TaskCtx interface {
	Assemble(ctx BeamerCtx) (pipe.Pipe, error)
}

type BeamerCtx struct {
	Tasks   []TaskCtx
	KVStore interface{}
	On      string
}

//############################################################################
func (p BeamerCtx) Expand(val string) (interface{}, error) {
	m := p.KVStore.(map[string]interface{})
	tmpl, err := template.New("*").Parse(val)
	if err != nil {
		return nil, err
	}

	buf := bytes.Buffer{}
	wr := bufio.NewWriter(&buf)

	if tmpl.Execute(wr, m); err != nil {
		return nil, err
	}

	key := string(buf.Bytes())
	return key, nil
}
