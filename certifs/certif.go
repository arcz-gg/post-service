package certifs

import (
	"path/filepath"
	"runtime"
)

var bPath string;

func _init() {
	_, currFile, _, _ :=runtime.Caller(0);
	bPath = filepath.Dir(currFile);
}

func Path(rel string) string {

	if filepath.IsAbs(rel) {
		return rel
	}

	return filepath.Join(bPath, rel)
}
