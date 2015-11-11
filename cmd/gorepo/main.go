package main

import (
	"fmt"
	"os"
	"path"
	"strings"
)

var repo_commands = make(map[string]func() int)
var repo_aliases = make(map[string][]string)

func main() {
	var cmd, arg0 string
	var args []string

	// cmd via arg0 using repo- or gorepo- prefixes
	arg0 = path.Base(os.Args[0])
	args = os.Args[1:]

	if strings.HasPrefix(cmd, "repo-") {
		cmd = arg0[5:]
	} else if strings.HasPrefix(cmd, "gorepo-") {
		cmd = arg0[7:]
	} else if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		// cmd via first argument, no prefix and not starting with a dash
		cmd = args[0]
		args = args[1:]
	} else {
		cmd = ""
	}

	// is cmd an alias?
	if alias, ok := repo_aliases[cmd]; ok {
		cmd = alias[0]
		args = append(alias[1:], args...)
	}

	if cmd != "" {
		// is cmd a valid command?
		if f, ok := repo_commands[cmd]; ok {
			arg0 = fmt.Sprintf("repo-%s", cmd)
			os.Args = append([]string{arg0}, args...)
			os.Exit(f())
		} else {
			fmt.Fprintf(os.Stderr, "%s: %s: Invalid command.\n", arg0, cmd)
		}
	} else {
		fmt.Fprintf(os.Stderr, "%s: No command requested\n", arg0)
	}

	// help
	fmt.Fprintf(os.Stderr, "Usage: %s <command> [options] arguments...\n", arg0)
	os.Exit(1)
}
