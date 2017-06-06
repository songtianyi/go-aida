# go-aida
go-aida是[wechat-go](http://github.com/songtianyi/wechat-go)的示例项目，包含Restful API和web页面两部分，web页面可以方便终端用户使用并管理wechat-go插件，API用来集成wechat-go和页面交互。

## 获取代码并运行
```
mkdir -p $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/net.git

cd $GOPATH/src/github.com/songtianyi/
git clone https://github.com/songtianyi/go-aida

cd go-aida/restful
go get ./...
go build .
./restful
```
## Restful API文档
domain http://your.domain:8080

#### /create

| /create| 创建一个机器人实例|
|------| ------ |
| **HEADER** ||
|||
| **PARAMS**||
|||


_Request_
```
GET /create
```
_Response_
```
200 OK
8c30a4e9-e949-4d10-b6d6-ef7b60e3af88
```

#### /status
| /status| 获取登录状态|
|------| ------ |
| **HEADER** ||
|||
| **PARAMS**||
|uuid|该session的uuid|

_Request_
```
GET /status?uuid=8c30a4e9-e949-4d10-b6d6-ef7b60e3af88
```
_Response_
```
200 OK
{
	"status": "CREATED",
	"qrcode": "../public/qrcode/wd_vvLuDWQ==.jpg",
	"plugins": {
		"laosj": true,
		"gifer": false
	},
	"startTime": 1496749513,
}
```
|status|意义|
|------| ------ |
|CREATED|等待用户扫码，此时已拿到二维码|

#### /enable

| /enable| 开启某个插件|
|------| ------ |
| **HEADER** ||
|||
| **PARAMS**||
|uuid|该session的uuid|
|name|插件名 eg. gifer|

_Request_
```
PATCH /enable?uuid=8c30a4e9-e949-4d10-b6d6-ef7b60e3af88&name=gifer
```

_Response_
```
200 OK
```

#### /disable

| /disable| 关闭某个插件|
|------| ------ |
| **HEADER** ||
|||
| **PARAMS**||
|uuid|该session的uuid|
|name|插件名 eg. gifer|


_Request_
```
PATCH /disable?uuid=8c30a4e9-e949-4d10-b6d6-ef7b60e3af88&name=gifer
```
_Response_
```
200 OK
```
