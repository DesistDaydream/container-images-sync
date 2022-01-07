package main

import (
	"os"

	"github.com/DesistDaydream/container-images-sync/pkg/flags"
	"github.com/DesistDaydream/container-images-sync/pkg/handle"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func LogInit(level, file string) error {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	le, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(le)

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
	logLevel := pflag.String("log-level", "info", "The logging level:[debug, info, warn, error, fatal]")
	logFile := pflag.String("log-output", "", "the file which log to, default stdout")
	imagesSyncFlags := &flags.ImagesSyncFlags{}
	imagesSyncFlags.AddFlags()
	pflag.Parse()

	// 初始化日志
	if err := LogInit(*logLevel, *logFile); err != nil {
		logrus.Fatal(errors.Wrap(err, "set log level error"))
	}
	handle.Handle(*imagesSyncFlags)
}
