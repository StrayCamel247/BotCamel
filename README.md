
## 聊天机器人-BotCamel
> 基于go/Mirai开发，扩展使用MiraiGo框架实现
## 项目启动
- 配置机器人qq账号密码
  复制`application_default_fmt.yaml`文件并更名为`application.yaml`，在文件中指定位置填写账号密码
  启动命令:`go run main.go`
## REFERENCE
- [敲代码的小小柒](https://www.bilibili.com/read/cv6926015/)
- [MiraiGo-Template](https://github.com/StrayCamel247/BotCamel/apps)
- [MiraiGo](https://github.com/Mrs4s/MiraiGo)
- ...若有遗漏请提醒
## 需求文档
| 模块 | 需求 | api/处理逻辑 | 需求状态 | 备注 |
|:------:|:------:|:------:|:------:|:------:|
| QQ机器人 | 按照文本一一对应回复 | 查询配置文件实现 | 完成 | 配置文件可进行扩充 |
| QQ-群组 | 判断是否@自己 | https://github.com/mamoe/mirai/blob/dev/docs/Messages.md | 完成 |  |
| QQ机器人 | 天气查询 |  |  |  |
| QQ-群组 | 命运2周报查询 |  |  |  |


<!-- ```
go mod
The commands are:
  download    download modules to local cache (下载依赖的module到本地cache))
  edit        edit go.mod from tools or scripts (编辑go.mod文件)
  graph       print module requirement graph (打印模块依赖图))
  init        initialize new module in current directory (再当前文件夹下初始化一个新的module, 创建go.mod文件))
  tidy        add missing and remove unused modules (增加丢失的module，去掉未用的module)
  vendor      make vendored copy of dependencies (将依赖复制到vendor下)
  verify      verify dependencies have expected content (校验依赖)
  why         explain why packages or modules are needed (解释为什么需要依赖)
``` -->