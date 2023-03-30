package shared_test

import (
	"abix360/shared"
	"strings"
	"testing"
)

func TestGetRootPath(t *testing.T) {
	rootPath := shared.GetRootPath()
	if !strings.Contains(rootPath, "cerbero") {
		t.Errorf("variable projectDirName tiene un valor diferente a cerbero")
	}
}
