package handle

import (
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/sirupsen/logrus"
)

func handleImages(srcImages, destImages, newRegistry string) (v1.Image, name.Reference, error) {
	var newDestImages string

	// 获取原镜像
	srcRef, err := name.ParseReference(srcImages)
	if err != nil {
		return nil, nil, err
	}

	srcImg, err := remote.Image(srcRef)
	if err != nil {
		return nil, nil, err
	}

	// fmt.Println(srcRef)
	// fmt.Println(srcRef.Context().Registry)
	// fmt.Println(srcRef.Context().Name())
	// fmt.Println(srcRef.Context().RepositoryStr())
	// fmt.Println(srcRef.Context().Scheme())
	// fmt.Println(srcRef.Context().Tag(srcRef.Identifier()))

	// 处理原镜像名称
	if destImages == "" {
		imageRepository := srcRef.Context().RepositoryStr()
		imageTag := srcRef.Identifier()
		newDestImages = newRegistry + "/" + imageRepository + ":" + imageTag
	} else {
		newDestImages = destImages
	}
	logrus.Debug("处理后的待推送镜像为:", newDestImages)

	// 指定镜像推送的目标
	destRef, err := name.ParseReference(newDestImages)
	if err != nil {
		return nil, nil, err
	}

	return srcImg, destRef, nil
}
