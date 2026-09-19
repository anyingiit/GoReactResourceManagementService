[English](README.md) · **简体中文**

> 英文版是规范版本。本页与 [README.md](README.md) 不一致时，以英文版为准。

<!-- translation-of: README.md sha256:be3e0ad4123e3a4f -->

<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# GoReactResourceManagementService

A Gin and GORM HTTP API that authenticates admins, users, and remote clients with separate token schemes, then hands each registered client a queue of tasks to run against configured web services and records the results in MySQL.

[![CI](https://github.com/anyingiit/GoReactResourceManagementService/actions/workflows/ci.yml/badge.svg)](https://github.com/anyingiit/GoReactResourceManagementService/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/anyingiit/GoReactResourceManagementService)](LICENSE)

[报告问题](https://github.com/anyingiit/GoReactResourceManagementService/issues/new?template=bug_report.yml) · [提出需求](https://github.com/anyingiit/GoReactResourceManagementService/issues/new?template=feature_request.yml)

<details>
  <summary>目录</summary>
  <ol>
    <li><a href="#about-the-project">关于本项目</a></li>
    <li><a href="#getting-started">开始使用</a></li>
    <li><a href="#usage">用法</a></li>
    <li><a href="#contributing">参与贡献</a></li>
    <li><a href="#license">许可证</a></li>
    <li><a href="#contact">联系方式</a></li>
  </ol>
</details>

## 关于本项目

GoReactResourceManagementService 只是 Go 后端本身：一个 Gin HTTP API（`main.go`），按角色对路由分组——public、user、admin、superadmin 和 client（`router/v1/init.go`）——尽管项目名里带有 "React"，这个仓库里并没有 React 前端代码。它用两套互不相同的机制来约束这些角色：面向管理员和普通用户的 JWT Bearer 令牌，以及面向远程客户端的按会话签发的 UUID 请求头（`middleware/auth.go`）。其核心数据模型是任务队列：每个已注册的客户端都会被分配一串针对某个已配置服务的任务，并通过 API 回报成功或失败，数据经由 GORM 保存在 MySQL 中（`models/task_queue.go`）。

计划中的功能与已知问题，见 [open issues](https://github.com/anyingiit/GoReactResourceManagementService/issues)。

## 开始使用

### 环境要求

- Go 1.20 或更高版本，即 `go.mod` 中固定的版本
- 一个可访问的 MySQL 服务；`docker-compose.yml` 提供了一个，监听 3306 端口
- 一份手写的 `config/config.yml`，字段需与 `structs/project_config.go` 一致——本仓库并未附带示例文件，缺少它时 `main.go` 会直接 panic

### 安装

```sh
git clone https://github.com/anyingiit/GoReactResourceManagementService.git
cd GoReactResourceManagementService

# main.go 会读取 config/config.yml，但这个文件并不在仓库中；
# 请先按照 structs/project_config.go 声明的 database、environments、
# server、token 字段创建它
mkdir -p config && $EDITOR config/config.yml

go mod download
go build -o main .
```

## 用法

`config/config.yml` 就绪后，可直接运行编译出的可执行文件：

```sh
./main
# Server is running on <server.local_ip>:<server.local_port>，具体值见 config/config.yml
```

或者一步启动 API、它的 MySQL 数据库以及一个 Adminer 管理界面：

```sh
docker compose up -d
```

`docker-compose.yml` 把 API 发布在 8080 端口，MySQL 发布在 3306 端口。

## 参与贡献

欢迎参与。[CONTRIBUTING.md](CONTRIBUTING.md) 说明如何提交 issue 或 pull request，[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) 说明对所有参与者的行为要求。

请不要在公开的 issue 或 pull request 中报告安全问题。[SECURITY.md](SECURITY.md) 说明了私下报告的方式。

## 许可证

以 MIT 许可证分发。详见 [LICENSE](LICENSE)。

## 联系方式

项目地址：[https://github.com/anyingiit/GoReactResourceManagementService](https://github.com/anyingiit/GoReactResourceManagementService)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
