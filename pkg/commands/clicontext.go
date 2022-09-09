package commands

import "os/user"

type CLIContext struct {
	Debug bool
	User  *user.User
}
