package config

type ClientControl struct {
	DisableBWM              bool
	DisableDownloadBWM      bool
	serverTimeDelta         int
	clientHost              string
	clientPort              int
	throttle_bytes          int
	disklimit_bytes         int64
	diskremaining_bytes     int64
	fileSystemBlocksize     int64
	rescanCache             bool
	verifyCache             bool
	useLessMemory           bool
	disableLogs             bool
	disableBWM              bool
	disableDownloadBWM      bool
	disableIPOriginCheck    bool
	disableFloodControl     bool
	skipFreeSpaceCheck      bool
	overrideConns           int
	maxAllowedFileSize      int
	staticRanges            map[string]int
	currentStaticRangeCount int
	cacheDirPath            string
	tempDirPath             string
	dataDirPath             string
	logDirPath              string
	downloadDirPath         string
	flushLogs               bool
	proxyHost               string
	proxyPort               int
	proxyType               string
	cacheSize               int
	useProxy                bool
}
