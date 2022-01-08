package handle

import (
	"strings"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/sirupsen/logrus"
)

type ImagesReference struct {
	source      string
	destination string
}

func HandleImages(imgRef ImagesReference, newRegistry string) (v1.Image, name.Reference, error) {
	var newDestImages string

	// 获取原镜像
	srcRef, err := name.ParseReference(imgRef.source)
	if err != nil {
		return nil, nil, err
	}

	srcImg, err := remote.Image(srcRef)
	if err != nil {
		return nil, nil, err
	}

	// 处理原镜像名称
	if imgRef.destination == "" {
		newRepository := handleRepository(srcRef)
		newTag := srcRef.Identifier()
		newDestImages = newRegistry + "/" + newRepository + ":" + newTag
	} else {
		newDestImages = imgRef.destination
	}
	logrus.Debug("处理后的待推送镜像为:", newDestImages)

	// 指定镜像推送的目标
	destRef, err := name.ParseReference(newDestImages)
	if err != nil {
		return nil, nil, err
	}

	return srcImg, destRef, nil
}

func handleRepository(srcRef name.Reference) string {
	imageRepository := srcRef.Context().RepositoryStr()
	// 有的镜像具有 Namespace，去要去掉
	slice := strings.SplitN(imageRepository, "/", 2)
	if len(slice) == 2 {
		imageRepository = slice[1]
	}
	return imageRepository
}
