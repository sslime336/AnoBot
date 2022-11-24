<div align="center"><img src="Bot.jpg" width="200" height="200" alt="图片显示不粗来力-w-"></div>

<div align="center">

# AnoBot

✨基于 [MiraiGo](https://github.com/Mrs4s/MiraiGo) ，参考 [go-cqhttp](https://github.com/Mrs4s/go-cqhttp) 的简单 QQ Bot 框架✨

</div>

---

<p align="center">本仓库所有代码仅作学习使用，不得用于其他任何途径</p>

---

基于 Go 1.19

## 特点

简单易用，数据分离

## 项目结构

- bot/ 基于 MiraiGo，参考 go-cqhttp 简单封装的 QQBot
  - export/ 通过 `import . "github.com/sslime336/awbot/bot/export"` 导入到模块中使用的机器人实体和添加定时任务的简单方法 `AddSchedule`
- config/ master，resources 等配置，负责从项目根目录读取 `bot.yaml` 中的相关配置
- logging/ 全局日志，默认是 [zap](https://github.com/uber-go/zap) 的 Dev
- modules/ Bot 的模块，可以自定义添加
  - libs 具体功能实现，可为外层模块提供支持
- res/ 数据常量
  - resources.json 该文件会在 Bot 初始化时读入到内存，参考 Android 中资源对象 R 的实现
- utils/ 简单封装了错误处理等逻辑，减少模板化代码
- main.go 启动机器人
- bot.yaml 机器人的配置文件

## 食用方法

1. 在项目根目录新建 `bot.yaml` 并参考 [bot_demo.yaml](bot_demo.yaml) 中的内容填写相关配置

2. 在 res/ 目录下参考 [resources_demo.json](res/resources_demo.json) 新建 `resources.json` 并填写相关内容

3. 在 modules/ 中编写您自己的模块，您可在 libs/ 中放置一些功能较为通用的模块，这些模块将被外层的模块导入并使用，不建议直接从这些模块中导出相关功能

4. 将您写好的模块注册到 [modules/mod.go](modules/mod.go)，具体的模块实现可参考 modules/ 中已有的模块

5. 运行 Bot，第一次登录需要扫码，二维码将生成在项目根目录 `qrcode.png`，需提前做好准备，之后不再需要扫码登录。您可通过在项目根目录运行 `make run` 或 `go run .` 来简单运行 Bot

## 有关 `make`

**make** - 默认指构建，不运行

- release 移除了所有符号表和可供调试的信息
- run 构建完成后运行
- clean 清理构建生成的可执行文件，以及设备信息文件
- *clean_all* 清理所有生成文件，包括登录相关的 token (清理后需要重新扫码登录)
