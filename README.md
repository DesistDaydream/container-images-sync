# image-syncer
使用 GitHub 的 Actions 功能定期同步一些海外镜像到国内可以访问的仓库中，同步时使用阿里的 [aliyun image-syncer](https://github.com/AliyunContainerService/image-syncer) 工具


在 google 的 [cloudshell](https://console.cloud.google.com/cloudshell) 可以查看 k8s.gcr.io 镜像的信息，示例命令如下：
```
gcloud container images list --repository=k8s.gcr.io/metrics-server
gcloud container images list-tags k8s.gcr.io/metrics-server/metrics-server
```