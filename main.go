package main

import (
	"github.com/mxuanp/go-utils/xfmt"

	"io/ioutil"
	"os"
)

func main() {
	path := "."
	if len(os.Args) != 1 {
		path = os.Args[1]
	}

	xfmt.PrintlnMsg(path, xfmt.Blue)
	println()
	tree(path, 1)
}

func tree(path string, dep int) {
	infos, err := ioutil.ReadDir(path)

	if err != nil {
		xfmt.PrintMsg(err.Error(), xfmt.Red)
		os.Exit(1)
	}

	for _, info := range infos {
		if info.IsDir() {
			xfmt.PrintMsg("|", xfmt.White)
			for t := dep; t >= 0; t-- {
				xfmt.PrintMsg("--", xfmt.White)
			}
			xfmt.PrintMsg(info.Name(), xfmt.Blue)
			xfmt.PrintlnMsg(" \\", xfmt.Blue)
			println()
			tree(info.Name(), dep+2)
			continue
		}

		xfmt.PrintMsg("|", xfmt.White)
		for t := dep; t >= 0; t-- {
			xfmt.PrintMsg("--", xfmt.White)
		}
		xfmt.PrintlnMsg(info.Name(), xfmt.White)
		println()
	}
}
