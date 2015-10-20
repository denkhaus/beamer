package main

import "github.com/denkhaus/beamer/tasks"

func main() {

	Beamer.Set("second_dir", "/home/denkhaus/test2")

	seq := Beamer.Sequence(
		tasks.Mkdir("/home/denkhaus/test1").
			Descr("create first directory"),
		tasks.Mkdir("{second_dir}").
			Descr("create second directory").,

		tasks.Copy.From("").To(""),
	)

	seq.Run()
}
