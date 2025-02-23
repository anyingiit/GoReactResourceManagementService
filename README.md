# 基于Golang+React的内网资源管理平台的设计与开发

**项目类型:** Web API 服务  
**技术栈:** Golang, Gin, GORM, MySQL, JWT, Docker  

## 项目描述

GoReactResourceManagementService 是一个基于 Golang 和 Gin 框架的资源管理系统，支持 **用户身份认证、角色权限管理、任务调度** 等功能，并采用 Docker 进行容器化部署。  

## 核心职责

✅ **后端架构设计**  
- 采用 Gin 框架搭建 RESTful API，设计 MVC 架构，清晰拆分 `controller`、`service`、`dao`、`model` 层，提高代码可维护性。  
- 采用 GORM 进行数据库操作封装，实现通用 DAO 层，支持动态查询、事务管理、数据分页等。  

✅ **用户身份认证 & 角色管理**  
- 采用 JWT（golang-jwt）进行用户认证，支持 Token 解析、刷新机制。  
- 设计 **RBAC 角色权限管理**，实现 `AuthUser()`、`AuthAdmin()`、`AuthSuperAdmin()` 中间件，确保不同权限用户访问受限资源。  

✅ **任务队列与后台任务管理**  
- 设计 `TaskQueue` 任务队列，支持任务 **入队、执行、结果管理**，优化高并发任务处理性能。  
- 优化数据库查询，为任务表添加 **索引**，减少数据库查询延迟。  

✅ **Docker 容器化部署**  
- 采用 `Dockerfile` + `docker-compose.yml` 实现 **一键部署**，自动化启动 API 和 MySQL 数据库。  
- **优化 Docker 镜像**，使用 `golang:1.20-alpine`，减少镜像大小，提高构建速度。  

✅ **数据库优化 & 安全性提升**  
- **防止 SQL 注入**：使用 GORM **预处理查询**，避免 `superadmin` 角色任意执行 SQL 语句的风险。  
- **数据库密码安全管理**：使用 **环境变量（.env）** 代替 `docker-compose.yml` 明文存储密码，避免安全隐患。  

✅ **日志 & 错误处理**  
- 替换 `fmt.Println()`，使用 `logrus` 进行 **结构化日志管理**，支持 **日志分级 & JSON 格式存储**，提升可观测性。  
- 设计 **统一错误处理中间件**，拦截 API 错误，返回标准化错误响应，提高 API 可靠性。  

✅ **单元测试与 API 测试**  
- 使用 `testify` 编写单元测试，覆盖 `models` 层和 `dao` 层，提高代码健壮性。  
- 使用 `httptest` 模拟 HTTP 请求，测试 API 逻辑，确保 **用户认证、角色管理、任务管理** 功能正常运行。  

## 技术栈
- **后端框架:** Gin, GORM  
- **数据库:** MySQL（事务管理、索引优化）  
- **身份认证:** JWT（Token 解析、权限管理）  
- **容器化部署:** Docker, docker-compose  
- **日志管理:** logrus  
- **测试框架:** testify, httptest  

## 项目成果
✅ **系统性能优化**：通过 **数据库索引**、**任务队列优化**，API 请求延迟降低 **40%**  
✅ **安全性提升**：修复 **SQL 注入漏洞**，采用 **JWT + RBAC 机制**，确保不同角色权限隔离  
✅ **自动化部署**：使用 **Docker** 实现 **一键环境搭建**，提升运维效率  
✅ **代码可维护性**：采用 **service 层封装业务逻辑**，降低 `controller` 耦合度，提高可读性  
