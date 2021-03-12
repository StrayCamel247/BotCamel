
## 聊天机器人-BotCamel_V1.0
> 基于/重构[go-cqhttp](https://github.com/Mrs4s/go-cqhttp/)实现-觉得有意思的话就[star]()一下吧~

> [BotCamel_V0](https://github.com/StrayCamel247/BotCamel/tree/dev-mirai)版本为基于mirai实现
## [需求文档](./PRD.MD)
> 欢迎任何人-领任务-[fork]()-提pr请求

任务领取可发送到aboyinsky@ouotlook.com邮箱
### 需求概览
## 项目启动
> 输入自己的账号密码可以构建自己的机器人哦
- windows
  - 安装go（很简单百度就会了）
  - 配置机器人qq账号密码
  - 启动命令:`go run main.go` 若没有配置文件会生成一个配置文件
  - `config.hjson`生成或者已存在
  - `apps\base_default.yaml`配置对话
  - `apps\baseapis`调用的接口
  - 若启动报错，尝试删除`go.mod`, `go.sum`文件，并运行`go mod init github.com/StrayCamel247/BotCamel`，再运行`go run main.go`启动
  - 若还是有问题请issue
  - 编译`go build`后点击exe文件即可后台运行
## REFERENCE

## DEMO

- version 1.1

  ![qq群聊演示V1.0](./media/motherfucker_asskisser.gif)

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