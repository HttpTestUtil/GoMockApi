# 关于 mock-api-go

该项目用于快速生成方便前端调试的 mock 服务器功能

## 关于服务器配置文件说明

- 配置文件默认在如下目录

` config/config.json `

- 关于配置项的说明

|key|type|comment|
|:---|:---|:---|
|listen_port|int|关于启动服务的时候，监听的文件类型|

（暂时只有一个配置项）

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

- 将本项目作为全局的 cli 命令，增加针对参数的支持

- 增加有限的随机生成数据的规则

- 增加有限的鉴权规则，方便部署在公网上进行测试

- 增加返回数据类型选项，字符串

- 增加返回数据类型选项，文件

- 制作相关工具，辅助生成和处理 api 文件。比如从 postman 迁移，比如图形化界面

## linsen

MIT
