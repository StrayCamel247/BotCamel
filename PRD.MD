# PRD-需求文档
> 欢饮提需求和ISSUE哦🙌

> 星级表示优先度-越多约优先-最多五颗星
## QQ机器人
### 基础功能
| 需求 | api/处理逻辑 | 需求状态 | 备注 |开发人 |
|:------:|:------:|:------:|:------:|:------:|
| 并发处理的实现 |  |  |  | |

### 指令表
### QQ私聊功能
| 需求 | api/处理逻辑 | 需求状态 | 备注 |开发人 |
|:------:|:------:|:------:|:------:|:------:|
| 按照文本一一对应回复 | 查询配置文件实现 | 完成 | 配置文件可进行扩充 |[StrayCamel247](https://github.com/StrayCamel247) |
| 沙雕APP-彩虹屁 | asskisser | 完成 |  |[StrayCamel247](https://github.com/StrayCamel247) |
| 沙雕APP-开光嘴 | motherfucker | 完成 |  |[StrayCamel247](https://github.com/StrayCamel247) |
| 沙雕APP-pyq文案生成 |  | 待认领 |  |
| 沙雕APP-毒鸡汤文案生成 |  | 待认领 |  |
| 菜单自动生成 | 想根据`apps\camel\command.go:CommandsStruct`结构体来进行自动生成，在被注释的`GenerateMenu`中尝试了一下，没成功，先待定 |  待认领|  |
| 按照文本一一对应回复 | 查询配置文件实现 | 完成 | 配置文件可进行扩充 |[StrayCamel247](https://github.com/StrayCamel247) |


### QQ群聊功能
| 需求 | api/处理逻辑 | 需求状态 | 备注 |开发人 |
|:------:|:------:|:------:|:------:|:------:|
| 判断是否@自己 | [消息结构体文档](https://github.com/mamoe/mirai/blob/dev/docs/Messages.md) | 完成 | 继承并使用QQ私聊功能 |[StrayCamel247](https://github.com/StrayCamel247) |
| 周报 | week | 完成 | [B站数据解析](https://api.bilibili.com/x/article/list/web/articles?id=175327&jsonp=jsonp) |[StrayCamel247](https://github.com/StrayCamel247) |[StrayCamel247](https://github.com/StrayCamel247) |
| 日报 | day | 完成 | [天阙提供](http://www.tianque.top/d2api/today/)-获取返回json 解析图片url |[StrayCamel247](https://github.com/StrayCamel247) |[StrayCamel247](https://github.com/StrayCamel247) |
| 试炼 | trial | 完成 | [B站数据解析](https://api.bilibili.com/x/article/list/web/articles?id=175690&jsonp=jsonp) |[StrayCamel247](https://github.com/StrayCamel247) |
| 老九 | nine | 完成 | [B站数据解析](https://api.bilibili.com/x/article/list/web/articles?id=175690&jsonp=jsonp) |[StrayCamel247](https://github.com/StrayCamel247) |
| 光尘商店 | dust | 完成 | [图片数据](https://cdn.jsdelivr.net/gh/azmiao/picture-bed/img/buy-13.jpg) |[StrayCamel247](https://github.com/StrayCamel247) |
| 棒极中文-英文-item-id数据存表 | 数据存表`data\sqlite3.db` | 完成 | [StrayCamel247](https://github.com/StrayCamel247) |
| 查看某把物品perk有哪些 |  | 完成 | [light-gg网页快照](https://www.light.gg/) |[StrayCamel247](https://github.com/StrayCamel247)|
| 自动加群 |  | 完成 |  func-GroReciveInviteHandler|[StrayCamel247](https://github.com/StrayCamel247)|
| item装备-武器查询 |  | 完成 | [light-gg网页快照](https://www.light.gg/) |[StrayCamel247](https://github.com/StrayCamel247)|
| npc信息查询|  | 完成 | 返回对信息的介绍 |[StrayCamel247](https://github.com/StrayCamel247)|
| perk信息查询|  | 完成 | 返回对信息的介绍 |[StrayCamel247](https://github.com/StrayCamel247)|
| skill光能技能查询|  | 完成 | 返回对信息的介绍 |[StrayCamel247](https://github.com/StrayCamel247)|
| 动态生成菜单|  |  |  ||
| ✨查pvp战绩 |  | 完成 | https://www.bungie.net/Platform/Destiny2/3/Account/xxx/Stats/|[StrayCamel247](https://github.com/StrayCamel247)|
| ✨查pve战绩 |  | 待认领 |  |
| ✨哪些perk属i于PVP或者PVE |  | 待认领 |  |
| ✨bot定时任务-每周三3.00 am触发 |  | 待认领 |  |
| ✨bot定时任务-每天3.00 am触发 |  | 待认领 |  |
| ✨可查询群文件-通过文件名来发送文件 | 触发-文件名-事件-发送文件 | 待认领 |  |

