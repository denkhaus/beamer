package tasks

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/denkhaus/beamer/common"
	"github.com/denkhaus/tcgl/applog"
	"github.com/juju/errors"

	"github.com/lann/builder"
	"gopkg.in/pipe.v2"
)

const REGEX_VAR_EXPAND = "^{{.*}}$"

type dirCtx struct {
	Directory   string
	FileMode    os.FileMode
	Description string
	Strict      bool
}

type dirBuilder builder.Builder

//############################################################################
func (b dirBuilder) FileMode(fm os.FileMode) dirBuilder {
	return builder.Set(b, "FileMode", fm).(dirBuilder)
}

//############################################################################
func (b dirBuilder) Descr(descr string) dirBuilder {
	return builder.Set(b, "Description", descr).(dirBuilder)
}

//############################################################################
func (b dirBuilder) Strict() dirBuilder {
	return builder.Set(b, "Strict", true).(dirBuilder)
}

//############################################################################
func (b dirBuilder) Assemble(ctx common.BeamerCtx) (pipe.Pipe, error) {
	d, ok := builder.Get(b, "Directory")
	if !ok {
		return nil, errors.New("mkdir: directory is undefined")
	}

	directory := d.(string)
	if ok, _ := regexp.MatchString(REGEX_VAR_EXPAND, directory); ok {
		d, err := ctx.Expand(directory)
		if err != nil {
			return nil, errors.Errorf("mkdir: expand error: %s", err)
		}
		directory = d.(string)
	}

	fm, ok := builder.Get(b, "FileMode")
	if !ok {
		return nil, errors.New("mkdir: filemode is undefined")
	}

	descr, ok := builder.Get(b, "Description")
	if !ok {
		descr = fmt.Sprintf("mkdir:: %q", directory)
	}

	strict := false
	if st, ok := builder.Get(b, "Strict"); ok {
		strict = st.(bool)
	}

	return func(s *pipe.State) error {
		applog.Infof("execute:: %s", descr)
		err := os.Mkdir(directory, fm.(os.FileMode))
		if !strict && strings.Contains(err.Error(), "file exists") {
			return nil
		}

		return err
	}, nil
}

//############################################################################
func Mkdir(directory string) dirBuilder {
	dir := builder.Register(dirBuilder{}, dirCtx{}).(dirBuilder)
	return builder.Set(dir, "Directory", directory).(dirBuilder)
}
