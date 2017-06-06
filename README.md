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
{
    "login": true,
    "plugins": {
        "laosj": true,
		"gifer": false
    },
	"startTime": 1496749513
}
```
