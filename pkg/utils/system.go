package utils

import (
	"os"
)

func UnsetEnvList(env ...string) {
	for _, envVar := range env {
		_ = os.Unsetenv(envVar)
	}
}
