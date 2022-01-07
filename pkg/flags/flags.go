package flags

import "github.com/spf13/pflag"

type ImagesSyncFlags struct {
	SrcUsername  string
	SrcPassword  string
	File         string
	DestRegistry string
}

func (isf *ImagesSyncFlags) AddFlags() {
	pflag.StringVarP(&isf.SrcUsername, "src-username", "", "", "用户名")
	pflag.StringVarP(&isf.SrcPassword, "src-password", "", "", "密码")
	pflag.StringVarP(&isf.File, "file", "f", "./images.yaml", "镜像列表文件")
	pflag.StringVarP(&isf.DestRegistry, "dest-registry", "", "", "目标仓库")
}
