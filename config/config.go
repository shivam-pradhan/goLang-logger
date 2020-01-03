package config

const TempDir = "/tmp"

// Logger paths
const (
	LogDir    = TempDir + "/logs"
	DebugFile = LogDir + "/debug.log"
	ErrorFile = LogDir + "/error.log"
)

// Log levels
const (
	LogDebug    = 0
	LogInfo     = 1
	LogWarn     = 2
	LogError    = 3
	LogDevPanic = 4
	LogPanic    = 5
	LogFatal    = 6
)

// Log levels map[int]string
var LogLevelStr = map[int]string{
	0: "LogDebug",
	1: "LogInfo",
	2: "LogWarn",
	3: "LogError",
	4: "LogDevPanic",
	5: "LogPanic",
	6: "LogFatal",
}

// current log level
var CurrentLogLevel int
