## Ray Shell

Most developers take command line shells like bash, zsh, fsh ...etc for granted (_including me_). But when I wanted to build one myself, that's when I realized that there is complex engineering right under my nose the whole time. Aspects of technology that goes into building a functional POSIX shell blew my mind.

- Process Lifecycle: Managing OS calls `fork()` and `exec()` calls to run external programs.
- System Architecture: Navigating the $PATH environment and managing file descriptors.
- Parsing Logic: Implementing a robust lexer to handle quoted strings, escaped characters and complex arguments.

### Features implemented so far

- [x] Basics (_Prompt_, _REPL_, Built in commands like _echo_, _exit_, _type_ & _OS system calls_)
- [x] Navigation (_pwd_, _cd_ & _~_)
- [x] Quoting (_'_, _"_ & _\\_)
- [x] Redirect (_input_, _output_ & _error_)
- [x] Pipelines
- [x] Autocompletion for both files and directories (_prefix matching via Trie, PATH scanning & Working directory_)
- [x] History (_in-session navigation with arrow keys_)
- [x] History persistence (_append/read/write to file_)

### Tools used
<img src="https://skillicons.dev/icons?i=go,neovim,git" />
