<img width="1645" height="1224" alt="image" src="https://github.com/user-attachments/assets/04ba932e-8051-4303-98bc-5d5e4c8f6c6c" /># HDU-KillCourse-LinuxDocker (Web 面板版)

[![Go Version](https://img.shields.io/badge/go-1.22+-blue.svg)](https://golang.org)
[![Docker](https://img.shields.io/badge/docker-compose-green.svg)](https://www.docker.com/)
[![Fork](https://img.shields.io/badge/fork-cr4n5%2FHDU--KillCourse-orange.svg)](https://github.com/cr4n5/HDU-KillCourse)

本项目是基于 [cr4n5/HDU-KillCourse](https://github.com/cr4n5/HDU-KillCourse) 深度重构与扩展的 **杭州电子科技大学多账号并发抢课 Docker Web 面板系统**。

在继承原版底层正方教务协议客户端（CAS 统一身份认证 / 钉钉扫码 / 新教务 RSA 登录 / 选退课接口）的基础上，增加了多账号常驻调度引擎、Web 控制台单页应用、Session 自动保活重登、DryRun 干跑保护及容器化一键部署方案。

---


## 🌟 核心特性

- 👥 **多账号并发调度**：支持多学号同时常驻挂机，每个账号拥有独立的会话生命周期、运行状态机与课程表优先级队列；
- 🔐 **双重保险登录链路**：
  - 优先复用本地 Cookie；
  - 会话失效时按设定优先级自动使用统一身份认证 CAS 或正方教务账密重登并回写 Cookie；
- 🛡️ **安全干跑模式 (DryRun)**：运行期支持开启干跑，可模拟完整的到点激活、会话刷新、课程余量查询与解析流程，不向服务器发起最终选退课提交；
- 📊 **现代化 Web 单页面板**：
  - 实时倒计时与服务器北京时间对齐；
  - 账号会话状态（有效/过期/来源）、课程执行结果（成功/失败/排队）一目了然；
  - 环形缓冲实时日志流展示与账号过滤；
  - 内置课程库缓存与一键拉取全校教学班信息（支持导出 Excel）；
- 🐳 **容器化部署与安全隔离**：
  - Docker Compose 一键拉起；
  - 支持 Caddy 提供反向代理、TLS 与 HTTP Basic Auth 访问鉴权。

---

## 📁 数据目录结构

运行期间数据持久化在挂载的 `./data` 目录：

```
data/
├── global.json           # 全局设置（定时开始时间/运行模式/轮询间隔/干跑开关/邮件通知）
├── accounts.json         # 账号队列（凭据、课程表、Cookie 与最新会话状态）
├── courses/{id}/         # 账号课程库缓存 (course.json) 与导出的任务落实 Excel
├── clientbody/{id}.json  # 选课控制参数缓存 (ClientBodyConfig)
└── log_files/            # 实时运行日志 (app.log、debug.log)
```

---

## 🚀 快速开始（Docker Compose 部署）

### 1. 克隆仓库
```bash
git clone https://github.com/lizhikeer/HDU-KillCourse-LinuxDocker.git
cd HDU-KillCourse-LinuxDocker
```

### 2. 配置反代访问凭据（可选，但推荐）
编辑 `deploy/caddy/Caddyfile`，配置你的反代端口或域名以及 Basic Auth 访问口令（可使用 `docker run --rm caddy caddy hash-password --plaintext "你的密码"` 生成哈希值）：

```caddy
:40000 {
    basic_auth {
        admin <生成的哈希值>
    }
    reverse_proxy grabber:40000
}
```

### 3. 一键启动
```bash
docker compose up -d --build
```

容器启动后：
- 本机调试端口：`http://127.0.0.1:40001`
- 公网反代端口：`https://<你的IP或域名>:40000`

### 4. 使用步骤
1. 打开网页管理面板，点击 **【＋ 添加账号】**，输入杭电学号与密码；
2. 点击账号行的 **【🧪 登录】**，测试验证登录与 Session 有效性；
3. 点击 **【✏ 编辑】** -> **【🔄 更新课程库】** 同步全校教学班信息；
4. 填入教学班编号（例如 `(2026-2027-1)-A0512040-06`），设定动作（选课 / 退课）及优先级排序；
5. 设置开始时间与运行模式，建议先勾选 **【开启干跑模式】** 进行验证，确认无误后正式挂机！

---

## 🛠️ 运维与调试

```bash
# 查看面板运行日志
docker compose logs -f grabber

# 重启容器服务（配置与会话无损保存于 ./data）
docker compose restart grabber

# 更新代码后重新构建
docker compose up -d --build
```

---

## 📜 鸣谢与开源许可

- 本项目基于 [cr4n5/HDU-KillCourse](https://github.com/cr4n5/HDU-KillCourse) 衍生开发；
- 遵循原项目开源许可证 [GPL-3.0 License](LICENSE)。仅供学习与技术研究交流使用，请勿用于商业用途。
