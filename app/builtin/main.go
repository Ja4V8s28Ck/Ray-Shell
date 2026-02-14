package builtin

type Builtin interface {
	Name() string
	Execute(shellArgs []string)
}

var CmdFuncMap = make(map[string]Builtin)

func init() {
	CmdFuncMap = map[string]Builtin{
		"cd":   Cd{},
		"echo": Echo{},
		"exit": Exit{},
		"pwd":  Pwd{},
		"type": Type{},
	}
}
