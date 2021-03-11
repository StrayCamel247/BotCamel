
## 聊天机器人-BotCamel
> 基于/重构[go-cqhttp](https://github.com/Mrs4s/go-cqhttp/)实现
## [需求文档](./PRD.MD)
- **急需** 邀请入群组自动同意
- qq群组-私聊基础功能开发
## 项目启动
- 配置机器人qq账号密码
  启动命令:`go run main.go` 若没有配置文件会生成一个配置文件
  `config.hjson`生成或者已存在
## REFERENCE

## DEMO
![qq群聊演示](./media/QQGOURPDEMO.gif)

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