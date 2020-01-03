# goLang-logger
Steps to run application on local:
1) clone this repo insside go/src
2) go run main.go

OutPut: -------------
2020-01-03T15:35:28.756+0530    INFO    ..... DEV LOG initialized successfully with CurrentLogLevel0(LogDebug)

3) hit curl:: curl localhost:8083/logLevel
you will get json response:  {"current_log_level":"LogDebug"}
console output:
2020-01-03T15:36:52.086+0530    DEBUG   reqID:clientreq_764595f3-d6c8-49d5-9844-adee16b266a8    .... Controller: GetCurrentLogLevel.
2020-01-03T15:36:52.086+0530    DEBUG   reqID:clientreq_764595f3-d6c8-49d5-9844-adee16b266a8    .... Current Log Level is :LogDebug

4) hit curl::  curl -X PUT localhost:8083/logLevel/1
you will get json response: {"current_log_level":"LogInfo"}
console output:
2020-01-03T15:38:04.184+0530    DEBUG   reqID:clientreq_f44b7aca-e182-4918-9f10-f496bb05a563    .... Controller: UpdateCurrentLogLevel.
2020-01-03T15:38:04.184+0530    DEBUG   reqID:clientreq_f44b7aca-e182-4918-9f10-f496bb05a563    .... Before Update, Current Log Level is :LogDebug
2020-01-03T15:38:04.184+0530    INFO    reqID:clientreq_f44b7aca-e182-4918-9f10-f496bb05a563    ....After Update, Current Log Level is :LogInfo
