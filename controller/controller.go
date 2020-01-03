package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"logger/config"
	"logger/utils"
	"net/http"
	"strconv"
)

// Hello ...
/**
 * just a controller to check /hello api
**/
func Hello(w http.ResponseWriter, r *http.Request) {
	reqID := utils.GetClientRequestUUID()
	utils.FileLogger("reqID:"+reqID+"	.... Controller: Hello.", config.CurrentLogLevel)
	jsonResult, err := json.Marshal("Hello, world")
	if err != nil {
		utils.FileLogger("reqID:"+reqID+"	.... Some error while processing json.", config.LogPanic)
	}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResult)
}

// UpdateCurrentLogLevel ...
/**
 * controller to update current log level
**/
func UpdateCurrentLogLevel(w http.ResponseWriter, r *http.Request) {
	reqID := utils.GetClientRequestUUID()
	pathParams := mux.Vars(r)
	utils.FileLogger("reqID:"+reqID+"	.... Controller: UpdateCurrentLogLevel.", config.CurrentLogLevel)
	utils.FileLogger("reqID:"+reqID+"	.... Before Update, Current Log Level is :" + config.LogLevelStr[config.CurrentLogLevel], config.CurrentLogLevel)
	currentLogLevel, inpErr  := strconv.Atoi(pathParams["logLevel"])
	if inpErr != nil {
		utils.FileLogger("reqID:"+reqID+"	.... Wrong Input.", config.LogPanic)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	config.CurrentLogLevel = currentLogLevel
	utils.FileLogger("reqID:"+reqID+"	....After Update, Current Log Level is :" + config.LogLevelStr[config.CurrentLogLevel], config.CurrentLogLevel)

	mapCurrentLogLevel := map[string]string {"current_log_level":config.LogLevelStr[config.CurrentLogLevel]}

	jsonResult, err := json.Marshal(mapCurrentLogLevel)
	if err != nil {
		utils.FileLogger("reqID:"+reqID+"	.... Some error while processing json.", config.LogPanic)
	}
	config.CurrentLogLevel,err = strconv.Atoi(pathParams["logLevel"])
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResult)
}

// GetCurrentLogLevel ...
/**
 * controller to get current log level
**/
func GetCurrentLogLevel(w http.ResponseWriter, r *http.Request) {
	reqID := utils.GetClientRequestUUID()
	utils.FileLogger("reqID:"+reqID+"	.... Controller: GetCurrentLogLevel.", config.CurrentLogLevel)
	mapCurrentLogLevel := map[string]string {"current_log_level":config.LogLevelStr[config.CurrentLogLevel]}
	jsonResult, err := json.Marshal(mapCurrentLogLevel)
	if err != nil {
		utils.FileLogger("reqID:"+reqID+"	.... Some error while processing json.", config.LogPanic)
	}
	utils.FileLogger("reqID:"+reqID+"	.... Current Log Level is :" + config.LogLevelStr[config.CurrentLogLevel], config.CurrentLogLevel)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResult)
}
