package constant

import (
	"os"
	P "path"
)

var Path *path

type path struct {
	CacheDir    string
	DataDir     string
	DownloadDir string
	LogDir      string
	TempDir     string
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir, _ = os.Getwd()
	}

	cacheDir := P.Join(homeDir, "cache")
	dataDir := P.Join(homeDir, "data")
	downloadDir := P.Join(homeDir, "download")
	logDir := P.Join(homeDir, "log")
	tmpDir := P.Join(homeDir, "tmp")
	Path = &path{
		CacheDir:    cacheDir,
		DataDir:     dataDir,
		DownloadDir: downloadDir,
		LogDir:      logDir,
		TempDir:     tmpDir,
	}

}
