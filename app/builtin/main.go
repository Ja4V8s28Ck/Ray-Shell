package builtin

import "io"

type ExecContext struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

type Builtin interface {
	Name() string
	Execute(shellArgs []string, ctx *ExecContext)
}

var CmdFuncMap = make(map[string]Builtin)

func init() {
	CmdFuncMap = map[string]Builtin{
		"cd":      Cd{},
		"echo":    Echo{},
		"exit":    Exit{},
		"history": History{},
		"pwd":     Pwd{},
		"type":    Type{},
	}

	buildTrie()
	BuildTrieLazy()
}
