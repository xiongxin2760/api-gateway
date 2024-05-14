package bootstrap

import (
	"api-gateway/library/resource"
	"context"
	"log"
	"path"
	"time"

	"api-gateway/pkg/format"
	"api-gateway/pkg/utils"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var LogPrefix string

// initLoggers 初始化logger
func initLoggers(_ context.Context) {
	// LogPrefix init
	LogPrefix = resource.Config.AppName + "_"

	var projectPath string
	if resource.Config.RunMode == "debug" {
		projectPath = path.Join(utils.GetProjectDir(), resource.Config.Logger.LogDir)
	} else {
		projectPath = path.Join(utils.GetParentDir(), resource.Config.Logger.LogDir)
	}

	// 初始化业务日志 logger 文件
	targetPath, err := utils.EnsurePwdDir(projectPath)
	if err != nil {
		panic(err)
	}

	serviceLogPathName := path.Join(targetPath,
		resource.Config.AppName+resource.Config.Logger.LogFileName)
	serviceLogFile, err := rotatelogs.New(
		serviceLogPathName+".%Y%m%d",
		rotatelogs.WithLinkName(serviceLogPathName),           // 默认 24 * time.Hour 分一次日志
		rotatelogs.WithMaxAge(14*time.Duration(24)*time.Hour), // 保留14天
	)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(serviceLogFile)                   // 设置日志输出到文件
	logrus.SetFormatter(&format.ServiceLogFormatter{}) // 日志格式
	logrus.SetReportCaller(true)                       // 日志是否打印堆栈
	log.Printf("Service Log Init Success")             // 写入日志内容
	tryRegisterCloser(serviceLogFile)

	// 初始化 gin access 日志 logger 文件
	ginAccessLogPathName := path.Join(targetPath, resource.Config.Logger.GinWebLogName)
	ginAccessLogFile, err := rotatelogs.New(
		ginAccessLogPathName+".%Y%m%d",
		rotatelogs.WithLinkName(ginAccessLogPathName),         // 默认 24 * time.Hour 分一次日志
		rotatelogs.WithMaxAge(14*time.Duration(24)*time.Hour), // 保留14天
	)
	if err != nil {
		panic(err)
	}
	resource.GinAccessLogger = logrus.New()                                // 实例化 gin access log
	resource.GinAccessLogger.SetOutput(ginAccessLogFile)                   // 设置输出
	resource.GinAccessLogger.SetFormatter(&format.GinAccessLogFormatter{}) // 设置日志格式
	log.Printf("Gin Access Log Init Success")
	tryRegisterCloser(ginAccessLogFile)

	// 初始化 日志 access 日志 logger 文件
	cronjobPathName := path.Join(targetPath, resource.Config.Logger.CronjobLogFileName)
	cronjobLogFile, err := rotatelogs.New(
		cronjobPathName+".%Y%m%d",
		rotatelogs.WithLinkName(cronjobPathName),              // 默认 24 * time.Hour 分一次日志
		rotatelogs.WithMaxAge(14*time.Duration(24)*time.Hour), // 保留14天
	)
	if err != nil {
		panic(err)
	}
	resource.CronjobLogger = logrus.New()                              // 实例化 gin access log
	resource.CronjobLogger.SetOutput(cronjobLogFile)                   // 设置输出
	resource.CronjobLogger.SetFormatter(&format.ServiceLogFormatter{}) // 设置日志格式
	resource.CronjobLogger.SetReportCaller(true)                       // 日志是否打印堆栈
	log.Printf("Crontab log init success")
	tryRegisterCloser(cronjobLogFile) // 写入日志内容

	// 初始化 recover 日志 logger 文件
	recoverLogPathName := path.Join(targetPath, resource.Config.Logger.RecoverLogFileName)
	recoverLogFile, err := rotatelogs.New(
		recoverLogPathName+".%Y%m%d",
		rotatelogs.WithLinkName(recoverLogPathName),           // 默认 24 * time.Hour 分一次日志
		rotatelogs.WithMaxAge(14*time.Duration(24)*time.Hour), // 保留14天
	)
	if err != nil {
		panic(err)
	}
	resource.RecoverLogger = logrus.New()                              // 实例化 recover log
	resource.RecoverLogger.SetOutput(recoverLogFile)                   // 设置输出
	resource.RecoverLogger.SetFormatter(&format.ServiceLogFormatter{}) // 设置日志格式
	resource.RecoverLogger.SetReportCaller(true)                       // 日志是否打印堆栈
	log.Printf("recover logger init success")
	tryRegisterCloser(recoverLogFile)
}
