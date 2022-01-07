package main

import (
	"fmt"
	"os"

	"github.com/DesistDaydream/container-images-sync/pkg/flags"
	"github.com/DesistDaydream/container-images-sync/pkg/handle"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/v1/remote"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

// LogInit 日志功能初始化，若指定了 log-output 命令行标志，则将日志写入到文件中
func LogInit(level, file, format string) error {
	switch format {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:   "2006-01-02 15:04:05",
			DisableTimestamp:  false,
			DisableHTMLEscape: false,
			DataKey:           "",
			// FieldMap:          map[logrus.fieldKey]string{},
			// CallerPrettyfier: func(*runtime.Frame) (string, string) {},
			PrettyPrint: false,
		})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	default:
		return fmt.Errorf("请指定正确的日志格式")
	}

	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(logLevel)

	if file != "" {
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			return err
		}
		logrus.SetOutput(f)
	}

	return nil
}

func main() {
	// 设置命令行标志
	logLevel := pflag.String("log-level", "info", "The logging level:[debug, info, warn, error, fatal]")
	logFile := pflag.String("log-output", "", "the file which log to, default stdout")
	logFormat := pflag.String("log-format", "text", "log format,one of: json|text")
	imagesSyncFlags := &flags.ImagesSyncFlags{}
	imagesSyncFlags.AddFlags()
	pflag.Parse()

	// 初始化日志
	if err := LogInit(*logLevel, *logFile, *logFormat); err != nil {
		logrus.Fatal(errors.Wrap(err, "set log level error"))
	}

	config := handle.NewConfig(imagesSyncFlags.File)

	// 实例化认证信息
	auth := authn.FromConfig(authn.AuthConfig{
		Username:      imagesSyncFlags.SrcUsername,
		Password:      imagesSyncFlags.SrcPassword,
		Auth:          "",
		IdentityToken: "",
		RegistryToken: "",
	})
	logrus.WithFields(logrus.Fields{
		"username": imagesSyncFlags.SrcUsername,
	}).Debug(auth)

	destOptions := remote.WithAuth(auth)

	// 开始同步镜像
	handle.Run(*config, destOptions)
}
