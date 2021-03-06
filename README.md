# 关于 mock-api-go

该项目用于快速生成方便前端调试的 mock 服务器功能

## 关于安装相关说明

* 现在的推荐安装方式是将该命令copy 到 /usr/bin/ 中作为一个全局命令使用。因为现在程序使用的配置文件目录，已经可以通过参数指定了。

## 可以指定的参数

|参数名|默认值|相关说明|
|:-----|:------|:------|
|port|8080|启动 http 服务的时候使用的接口|
|api-path|./api|关于Api文件的路径（仅支持目录）|

## 关于 API 配置文件的说明

- 程序会自动加载 api 目录下的所有 json 文件（暂时不支持递归加载

- json 数据格式示例

```json
    [{
        "route": "/ping",
        "method": "get",
        "response": {
            "message": "pong",
            "data": {
                "a": 1,
                "b": 2,
                "c": 3
            }
        }
    }]
```

- 相关参数说明

|key|type|comment|
|:---|:---|:---|
|route|string|触发该请求的路径|
|method|string|触发该请求的方法仅支持 get 和 post，其它类型会被忽略|
|response|json|作为返回值的 json 内容|

## Roadmap

- 将本项目作为全局的 cli 命令，增加针对参数的支持(check)

- 增加有限的随机生成数据的规则

- 增加有限的鉴权规则，方便部署在公网上进行测试

- 增加返回数据类型选项，字符串

- 增加返回数据类型选项，文件

- 制作相关工具，辅助生成和处理 api 文件。比如从 postman 迁移，比如图形化界面

## linsen

MIT
