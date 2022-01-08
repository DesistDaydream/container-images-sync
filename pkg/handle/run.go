package handle

import (
	"os"
	"sync"

	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type HandleTask struct {
	Success int
	Failed  int
}

type Config struct {
	ImagesList map[string]string `yaml:"images"`
}

func NewConfig(file string) *Config {
	var config Config

	imagesListFile, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(imagesListFile, &config.ImagesList)
	if err != nil {
		panic(err)
	}

	return &config
}

func Run(config Config, destOptions remote.Option, newRegistry string) {
	var wg sync.WaitGroup
	defer wg.Wait()

	for source, destination := range config.ImagesList {
		wg.Add(1)
		var imgRef ImagesReference

		imgRef.source = source
		imgRef.destination = destination
		logrus.WithFields(logrus.Fields{
			"src":  imgRef.source,
			"dest": imgRef.destination,
		}).Info("待处理镜像列表")

		go func(imgRef ImagesReference) {
			defer wg.Done()
			// 处理原镜像与目的镜像
			srcImg, destRef, err := HandleImages(imgRef, newRegistry)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"src":    imgRef.source,
					"dest":   imgRef.destination,
					"reason": err,
				}).Error("失败处理的镜像")
			} else if srcImg != nil && destRef != nil {
				logrus.WithFields(logrus.Fields{
					"src":  srcImg,
					"dest": destRef,
				}).Info("成功处理的镜像")

				// 使用 options 中的认证信息，将 img 推送到 ref 中
				// err = remote.Write(destRef, srcImg, destOptions)
				// if err != nil {
				// 	logrus.Error("镜像推送失败:", err)
				// }
			} else {
				logrus.Error("处理镜像未知错误")
			}
		}(imgRef)
	}
}
