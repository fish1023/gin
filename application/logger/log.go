package logger

import (
	"captcha/application/lib/helper"
	"os"
	"time"
)

// LEVEL_MAP log级别对应的文件名
var LEVEL_MAP = map[string]string{
	//调试
	"debug": "debug",
	//错误
	"error":   "error",
	"warn":    "error",
	"warning": "error",
	//统计
	"crit":    "stat",
	"stat":    "stat",
	"service": "service",
	//通用信息
	"notice":  "info",
	"info":    "info",
	"default": "info",
}

// LogFish log记录结构体
type LogFish struct {
	Cate string
}

// Debug debug log
func (l *LogFish) Debug(msg string) {
	l.appendLog(msg, "debug")
}

// Info info log
func (l *LogFish) Info(msg string) {
	l.appendLog(msg, "info")
}

// Error error log
func (l *LogFish) Error(msg string) {
	l.appendLog(msg, "error")
}

// Warn warn log
func (l *LogFish) Warn(msg string) {
	l.appendLog(msg, "warn")
}

// Stat stat log
func (l *LogFish) Stat(msg string) {
	l.appendLog(msg, "stat")
}

func (l *LogFish) appendLog(msg string, level string) {
	logFile := getLogPath(level)
	fd, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	fd_time := time.Now().Format("2006-01-02 15:04:05")
	hostname, _ := os.Hostname()
	msg = hostname + "\t" + fd_time + "\t" + level + "\t" + l.Cate + "\t" + msg + "\n"
	buf := []byte(msg)
	fd.Write(buf)
	fd.Close()
}

// getLogPath 获取级别对应的log文件
func getLogPath(level string) string {
	return helper.GetLogPath() + LEVEL_MAP[level]
}
