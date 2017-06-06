# go-aida
go-aida是[wechat-go](http://github.com/songtianyi/wechat-go)的示例项目，包括restful api和web页面两部分，web页面方便终端用户使用wechat-go默认插件或第三方插件的功能，restful api用来和wechat-go交互。

## Restful API文档
domain http://your.domain:8080

#### /qrcode
||||
|------| ------ | ------ |
| 描述 | 获取微信登录二维码 |

_Request_
```
GET /qrcode
```
_Response_
```
200 OK
binary body
```

#### /status
||||
|------| ------ | ------ |
| 描述 | 获取登录状态|

_Request_
```
GET /status
```
_Response_
```
200 OK
{
	"login": true,
	"plugins": {
		"laosj": true,
		"gifer": false
    },
	"startTime": 1496749513,
}
```

#### /enable
||||
|------| ------ | ------ |
| 描述 | 开启某个插件|
| name | 插件名 | eg. gifer|

_Request_
```
PATCH /enable?name=gifer
```

_Response_
```
200 OK
```

#### /disable
||||
|------| ------ | ------ |
| 描述 | 关闭某个插件|
| name | 插件名 | eg. gifer|

_Request_
```
PATCH /disable?name=gifer
```
_Response_
```
200 OK
```


