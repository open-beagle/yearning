# yearning

1. 请填写 conf.toml 相关配置信息。配置项与 Yearning1.0 基本相同
2. ./yearning install
3. ./yearning run

详细参数 请执行./Yearning --help 获得帮助信息

```bash
git remote add upstream git@github.com:cookieY/Yearning.git

git fetch upstream

git merge 2.3.5
```

## build

```bash
docker pull zhangsean/yearning:2.3.5 && \
docker tag zhangsean/yearning:2.3.5 registry.cn-qingdao.aliyuncs.com/wod/yearning:2.3.5 && \
docker push registry.cn-qingdao.aliyuncs.com/wod/yearning:2.3.5

docker run \
--rm \
-v $PWD/:/go/src/gitlab.wodcloud.com/cloud/yearning \
-v /go/pkg/:/go/pkg \
-w /go/src/gitlab.wodcloud.com/cloud/yearning \
-e PLUGIN_BINARY=yearning \
-e CI_WORKSPACE=/go/src/gitlab.wodcloud.com/cloud/yearning \
registry.cn-qingdao.aliyuncs.com/wod/devops-go-arch:1.16.7-alpine

# frontend
docker run \
-it \
--rm \
-v $PWD/:/go/src/gitlab.wodcloud.com/cloud/yearning \
-w /go/src/gitlab.wodcloud.com/cloud/yearning \
registry.cn-qingdao.aliyuncs.com/wod/yearning-gemini:v2.3.2 \
ash -c 'rm -rf ./dist/www && mkdir -p ./dist/www && cp -r /dist/* ./dist/www'
```
