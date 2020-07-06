# go实现简单生成二维码

## go build
```shell script
# 构建linux 64位环境运行
GOOS=linux GOARCH=amd64 go build -o linux_amd64_httpQrcodeServer httpQrcodeServer.go
# 构建macos 64位环境运行
GOOS=darwin GOARCH=amd64 go build -o darwin_amd64_httpQrcodeServer httpQrcodeServer.go

```

## 运行方式
1. linux or macos 其它可执行构建
./linux_amd64_httpQrcodeServer 9090 &

默认为9090端口，可省略

## 使用方式

GET /createQrcode?text={}&size={}

text二维码内容
size为图片大小，最大限制了512
encode=url时会进行urldecode处理

例如：
http://localhost:9999/createQrcode?text=https%3a%2f%2fwww.yuque.com%2fjnliok%2fpub%2fgolang&size=256&encode=url

