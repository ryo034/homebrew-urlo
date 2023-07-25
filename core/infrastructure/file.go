package infrastructure

import (
	"os"
	"path/filepath"
)

var (
	fileName         = "urls.json"
	dataDirName      = ".urlo"
	homeDir, _       = os.UserHomeDir()
	dataDir          = filepath.Join(homeDir, dataDirName)
	FileRelativePath = filepath.Join(dataDir, fileName)
)
