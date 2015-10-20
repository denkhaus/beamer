package main

import (
	"bitbucket.org/denkhaus/mirsvc/util"
	"github.com/denkhaus/beamer/tasks"
	"github.com/denkhaus/tcgl/applog"
)

func main() {

	var beamer = Beamer.
		Set("Second_dir", "/home/denkhaus/test2").
		Set("Loc1", "localhost").
		On("lulu")

	beamer = beamer.Sequence(
		tasks.Mkdir("/home/denkhaus/test1").FileMode(0777).Descr("create first directory"),
		tasks.Mkdir("{{.Second_dir}}").Descr("create directory by key/value"),
	//	//	//tasks.Copy.From("/home/denkhaus/test1/testfile").To("{{.Second_dir}}/testfile").Descr("Copy important file"),
	)

	if err := beamer.Run(); err != nil {
		applog.Errorf("runtime error: %s", err)
	}

	util.Inspect(beamer.Build())
}
