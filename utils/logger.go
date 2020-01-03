package utils

import (
	"fmt"
	"logger/config"
	"os"
	"strconv"

	"go.uber.org/zap"
)

// Logger ...
type LoggerStruct struct{
	Log *zap.Logger
}

var logger *LoggerStruct

// init ...
/**
 * initialize the log environment depending upon PROD/DEV env
 **/
func init()  {
	switch os.Getenv("ENV") {
	case "PROD":
		logger.Log, _ = zap.NewProduction()
		logger.Log.Info("..... PROD LOG initialized successfully.")
		break

	case "", "DEV":
		zapConfig := zap.NewDevelopmentConfig()
		zapConfig.DisableCaller = true
		if err := os.MkdirAll(config.LogDir, os.ModePerm); err != nil {
			panic("Creating debug log directory failed")
		}
		zapConfig.OutputPaths = append(zapConfig.OutputPaths, config.DebugFile)
		zapConfig.ErrorOutputPaths = append(zapConfig.ErrorOutputPaths, config.ErrorFile)
		log, err := zapConfig.Build()
		if err != nil {
			panic(fmt.Sprintf("..... DEV LOG initialization failed %v.", err))
		}
		logger = &LoggerStruct{Log:log}
		logger.Log.Info("..... DEV LOG initialized successfully with CurrentLogLevel" + strconv.Itoa(config.CurrentLogLevel) + "("+ config.LogLevelStr[config.CurrentLogLevel]+")")
		break
	}
}


func FileLogger(logString string, logLevel int) {
	switch logLevel {
	case config.LogDebug:
		logger.Log.Debug(logString)
		break
	case config.LogInfo:
		logger.Log.Info(logString)
		break
	case config.LogError:
		logger.Log.Error(logString)
		break
	case config.LogWarn:
		logger.Log.Warn(logString)
		break
	case config.LogDevPanic:
		logger.Log.DPanic(logString)
		break
	case config.LogPanic:
		logger.Log.Panic(logString)
		break
	case config.LogFatal:
		logger.Log.Fatal(logString)
		break

	default:
		logger.Log.Debug(logString)
		break
	}
}
