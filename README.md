# 功能

- 验证码服务

# 项目目录结构

```
  + application             // 代码逻辑
    |+ lib                  // 工具包
    |+ logger                 // log包
    |+ model                  // 数据交互
    |+ modules              // 模块
      |+img                 // 验证码模块
    |+ setting              // 配置启动
  + conf                      // 配置文件
    |- app.yaml             // 应用配置
    |- bjyt_pub_infra.ini   // qbus配置
  - Dockerfile
  + vendor                  // 第三方包
  - glide.lock              // 第三方包版本控制
  - glide.yaml              // 第三方包列表
  - main.go                 // 程序入口
```

# 代码发布流程

## 依赖

- [golang 1.8 版本](https://golang.org/dl/)
- [glide 包管理工具](https://github.com/Masterminds/glide)
- [Docker 1.8 以上 Client](https://www.docker.com/)

## 发布示例

1. 安装依赖&构建

```bash
  cd /path/项目名
  make glide
  make vendor
  make img
```

2. 提交到release分支&Hulk自动集成镜像

```bash
  git checkout release
  git merge master
  git add .
  git commit -m "publish new version"

```
3. Hulk更新应用镜像
