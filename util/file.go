package util

import (
	"os"
	"path/filepath"
)

var (
	homeDir, _       = os.UserHomeDir()
	dataDir          = filepath.Join(homeDir, ".urlo")
	FileRelativePath = filepath.Join(dataDir, "urls.json")
)
