package flags

import "github.com/spf13/pflag"

type ImagesSyncFlags struct {
	Username string
	Password string
	File     string
}

func (isf *ImagesSyncFlags) AddFlags() {
	pflag.StringVarP(&isf.Username, "username", "u", "", "用户名")
	pflag.StringVarP(&isf.Password, "password", "p", "", "密码")
	pflag.StringVarP(&isf.File, "file", "f", "./images.yaml", "镜像列表文件")
}
