# image-syncer
使用 GitHub 的 Actions 功能定期同步一些海外镜像到国内可以访问的仓库中，同步时使用阿里的 [aliyun image-syncer](https://github.com/AliyunContainerService/image-syncer) 工具



![手动触发工作流](/pictrues/手动触发工作流.png)


在 google 的 [cloudshell](https://console.cloud.google.com/cloudshell) 可以查看 k8s.gcr.io 镜像的信息，示例命令如下：
```bash
gcloud container images list --repository=k8s.gcr.io/metrics-server
gcloud container images list-tags k8s.gcr.io/metrics-server/metrics-server
```



```json
{
    "k8s.gcr.io/kube-apiserver": "docker.io/lchdzh/kube-apiserver",
    "k8s.gcr.io/kube-controller-manager": "docker.io/lchdzh/kube-controller-manager",
    "k8s.gcr.io/kube-scheduler": "docker.io/lchdzh/kube-scheduler",
    "k8s.gcr.io/kube-proxy": "docker.io/lchdzh/kube-proxy",
    "k8s.gcr.io/pause": "docker.io/lchdzh/pause",
    "k8s.gcr.io/etcd": "docker.io/lchdzh/etcd",
    "k8s.gcr.io/coredns": "docker.io/lchdzh/coredns",
    "k8s.gcr.io/ingress-nginx/controller": "docker.io/lchdzh/ingress-nginx-controller",
    "k8s.gcr.io/ingress-nginx/kube-webhook-certgen": "docker.io/lchdzh/kube-webhook-certgen",
    "k8s.gcr.io/kube-state-metrics/kube-state-metrics": "docker.io/lchdzh/kube-state-metrics",
    "k8s.gcr.io/metrics-server/metrics-server": "docker.io/lchdzh/metrics-server",
    "k8s.gcr.io/sig-storage/nfs-subdir-external-provisioner": "docker.io/lchdzh/nfs-subdir-external-provisioner",
    "quay.io/coreos/flannel": "docker.io/lchdzh/flannel",
    "quay.io/thanos/thanos": "docker.io/lchdzh/thanos"
}

```


```json
{
    "quay.io": {
        "username": "xxx",
        "password": "xxxxxxxxx",
        "insecure": true
    },
    "registry.cn-beijing.aliyuncs.com": {
        "username": "xxx",
        "password": "xxxxxxxxx"
    },
    "registry.hub.docker.com": {
        "username": "DOCKER_USERNAME",
        "password": "DOCKER_PASSWORD"
    },
    "quay.io/coreos": {
        "username": "abc",
        "password": "xxxxxxxxx",
        "insecure": true
    }
}
```



{
    "docker.io/kiwigrid/k8s-sidecar": "registry.cn-hangzhou.aliyuncs.com/dd_k8s/k8s-sidecar"
    // "k8s.gcr.io/kube-apiserver": "registry.cn-hangzhou.aliyuncs.com/dd_k8s/kube-apiserver",
    // "k8s.gcr.io/kube-controller-manager": "registry.cn-hangzhou.aliyuncs.com/dd_k8s/kube-controller-manager",
    // "k8s.gcr.io/kube-scheduler": "registry.cn-hangzhou.aliyuncs.com/dd_k8s/kube-scheduler",
    // "k8s.gcr.io/kube-proxy": "registry.cn-hangzhou.aliyuncs.com/dd_k8s/kube-proxy",
    // "k8s.gcr.io/pause": "registry.cn-hangzhou.aliyuncs.com/dd_k8s/pause",
    // "k8s.gcr.io/etcd": "registry.cn-hangzhou.aliyuncs.com/dd_k8s/etcd",
    // "k8s.gcr.io/coredns": "registry.cn-hangzhou.aliyuncs.com/dd_k8s/coredns",
    // "k8s.gcr.io/ingress-nginx/controller": "registry.cn-hangzhou.aliyuncs.com/dd_k8s/ingress-nginx-controller",
    // "k8s.gcr.io/ingress-nginx/kube-webhook-certgen": "registry.cn-hangzhou.aliyuncs.com/dd_k8s/kube-webhook-certgen",
    // "k8s.gcr.io/kube-state-metrics/kube-state-metrics": "registry.cn-hangzhou.aliyuncs.com/dd_k8s/kube-state-metrics",
    // "k8s.gcr.io/metrics-server/metrics-server": "registry.cn-hangzhou.aliyuncs.com/dd_k8s/metrics-server",
    // "k8s.gcr.io/sig-storage/nfs-subdir-external-provisioner": "registry.cn-hangzhou.aliyuncs.com/dd_k8s/nfs-subdir-external-provisioner",
    // "quay.io/coreos/flannel": "registry.cn-hangzhou.aliyuncs.com/dd_k8s/flannel",
    // "quay.io/thanos/thanos": "registry.cn-hangzhou.aliyuncs.com/dd_k8s/thanos"
}
