package handle

import (
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func handleImages(srcImages, destImages string) (v1.Image, name.Reference, error) {
	// 获取源镜像
	srcRef, err := name.ParseReference(srcImages)
	if err != nil {
		return nil, nil, err
	}

	srcImg, err := remote.Image(srcRef)
	if err != nil {
		return nil, nil, err
	}

	// 指定镜像推送的目标
	destRef, err := name.ParseReference(destImages)
	if err != nil {
		return nil, nil, err
	}

	return srcImg, destRef, nil
}
