package utils

import (
	"os"
	"os/user"
	"path/filepath"
)

func UnsetEnvList(env ...string) {
	for _, envVar := range env {
		_ = os.Unsetenv(envVar)
	}
}

func Expand(path string, currentUser *user.User) (string, error) {
	if len(path) == 0 || path[0] != '~' {
		return path, nil
	}

	return filepath.Join(currentUser.HomeDir, path[1:]), nil
}
