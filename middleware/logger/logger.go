package logger

import (
	"ClockInLite/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	logFilePath := config.ServerSetting.LogPath
	logFileName := config.ServerSetting.LogName

	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logWriter, err := rotatelogs.New(
		//分割后的文件名称
		fileName+".%Y%m%d.log",

		//生成软链，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(fileName),

		//设置最大保存时间（7天）
		rotatelogs.WithMaxAge(7*24*time.Hour),

		//设置日志切割时间间隔（1天）
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}, &logrus.JSONFormatter{})

	//添加Hook
	logrus.AddHook(lfHook)

	return func(c *gin.Context) {

		//开始时间
		startTime := time.Now()

		//处理请求
		c.Next()

		//结束时间
		endTime := time.Now()

		//执行时间
		latencyTime := endTime.Sub(startTime)

		//请求方式
		reqMethod := c.Request.Method

		//请求路由
		reqUri := c.Request.RequestURI

		//状态码
		statusCode := c.Writer.Status()

		//请求IP
		clientIP := c.ClientIP()

		//日志格式
		logger.WithFields(logrus.Fields{
			"StatusCode":    statusCode,
			"LatencyTime":   latencyTime,
			"ClientIp":      clientIP,
			"RequestMethod": reqMethod,
			"RequestUrl":    reqUri,
		}).Info()
	}
}

// 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
