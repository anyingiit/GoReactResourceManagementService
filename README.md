<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# GoReactResourceManagementService

A Gin and GORM HTTP API that authenticates admins, users, and remote clients with separate token schemes, then hands each registered client a queue of tasks to run against configured web services and records the results in MySQL.

**English** · [简体中文](README.zh-CN.md)

[![CI](https://github.com/anyingiit/GoReactResourceManagementService/actions/workflows/ci.yml/badge.svg)](https://github.com/anyingiit/GoReactResourceManagementService/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/anyingiit/GoReactResourceManagementService)](LICENSE)

[Report a bug](https://github.com/anyingiit/GoReactResourceManagementService/issues/new?template=bug_report.yml) · [Request a feature](https://github.com/anyingiit/GoReactResourceManagementService/issues/new?template=feature_request.yml)

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

## About The Project

GoReactResourceManagementService is the Go backend only: a Gin HTTP API (`main.go`) that groups its routes by role — public, user, admin, superadmin, and client (`router/v1/init.go`) — with no React frontend code in this repository despite the project name. It enforces those roles with two separate schemes: a JWT bearer token for admins and users, and a per-session UUID header for remote clients (`middleware/auth.go`). Its central data model is a task queue: each registered client is assigned a sequence of tasks against a configured web service and reports success or failure back through the API, backed by MySQL via GORM (`models/task_queue.go`).

See the [open issues](https://github.com/anyingiit/GoReactResourceManagementService/issues) for planned features and known issues.

## Getting Started

### Prerequisites

- Go 1.20 or newer, the version pinned in `go.mod`
- A reachable MySQL server; `docker-compose.yml` provides one, listening on port 3306
- A hand-written `config/config.yml` matching `structs/project_config.go` — the repository does not ship one, and `main.go` panics without it

### Installation

```sh
git clone https://github.com/anyingiit/GoReactResourceManagementService.git
cd GoReactResourceManagementService

# main.go reads config/config.yml, which is not part of this repository;
# create one with the database, environments, server and token fields that
# structs/project_config.go declares before going further
mkdir -p config && $EDITOR config/config.yml

go mod download
go build -o main .
```

## Usage

Once `config/config.yml` exists, run the binary directly:

```sh
./main
# Server is running on <server.local_ip>:<server.local_port>, as set in config/config.yml
```

Or bring up the API together with its MySQL database and an Adminer UI in one step:

```sh
docker compose up -d
```

`docker-compose.yml` publishes the API on port 8080 and MySQL on port 3306.

## Contributing

Contributions are welcome. Read [CONTRIBUTING.md](CONTRIBUTING.md) for how to open an issue or a pull request, and [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for the standards expected of everyone taking part.

Please do not report security issues in public issues or pull requests. [SECURITY.md](SECURITY.md) explains how to report them privately.

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for details.

## Contact

Project link: [https://github.com/anyingiit/GoReactResourceManagementService](https://github.com/anyingiit/GoReactResourceManagementService)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
