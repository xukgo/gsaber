package envUtil

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func GetProjectRootPathSkip(skip int) (string, error) {
	p, err := GetProjectRootByGoCmd()
	if err == nil {
		return p, nil
	}

	p, err = GetProjectRootByRuntimeSkip(skip + 2)
	if err == nil {
		return p, nil
	}
	return p, err
}
func GetProjectRootPath() (string, error) {
	p, err := GetProjectRootByGoCmd()
	if err == nil {
		return p, nil
	}

	p, err = GetProjectRootByRuntimeSkip(2)
	if err == nil {
		return p, nil
	}
	return p, err
}

func GetProjectRootByRuntimeSkip(skip int) (string, error) {
	_, filename, _, ok := runtime.Caller(skip)
	if !ok {
		return "", fmt.Errorf("cannot get invoker info")
	}

	dir := filepath.Dir(filename)
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod file notfound")
		}
		dir = parent
	}
}
func GetProjectRootByGoCmd() (string, error) {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}
