# 🎵  API

## 部署

- 极力推荐使用`docker compose`

  1. 安装`Docker`

  2. 命令行进入该项目目录下
     ```bash
     # 只需要目录中有 该项目中的 docker-compose.yml 文件即可
     docker compose up -d
     ```

     > 服务已可在本地 8080 端口访问了

- 本地部署 - 很麻烦

  1. 安装好 `go`、`node`

  2. 下载该项目和 [该项目依赖的网易接口](https://github.com/Binaryify/NeteaseCloudMusicApi)

  3. 修改 `config` 目录下的 `config.yaml`文件
     ```yaml
     # 原内容：
     ProxyAddr: http://netease_music_api:3000
     # ProxyAddr: http://127.0.0.1:3000
     
     #修改为：
     # ProxyAddr: http://netease_music_api:3000
     ProxyAddr: http://127.0.0.1:3000
     ```

  4. 运行 `NodeJs` 的接口项目
     ```bash
     npm install
     npm run start
     ```

  5. 运行本项目
     ```bash
     go run main.go
     ```

     > 服务已可在本地 8080 端口访问了

## 已完成API：

- [X] 歌单相关
  - [X] 默认歌单
  - [X] 推荐歌单
  - [X] 歌单详情
  - [X] 用于歌单分类的🏷️
  - [X] 通过🏷️获取相关歌单
- [X] 🔍相关
  - [X] 🔍提示
  - [X] 单曲🔍
  - [X] 歌手🔍
  - [X] 专辑🔍
  - [X] 歌单🔍
  - [X] MV🔍
- [X] 👨‍🎤相关
  - [X] 👩‍🎤相关🎵
  - [X] 👩‍🎤相关专辑
- [X] 🎵相关
  - [X] 🎵URL
  - [X] 🎵歌词
  - [X] 🎵MV
  - [X] 🎵详情
- [X] 排行榜相关
  - [X] 排行榜列表
  - [X] 排行榜详情

## 感谢

- [TS写的api接口](https://github.com/QiuYaohong/kuwoMusicApi.git)
- [非常出名的 NodeJs 版网易云接口代理](https://github.com/Binaryify/NeteaseCloudMusicApi)
