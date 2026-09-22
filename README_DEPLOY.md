# HDU 抢课面板 · 部署指南

多账号并发抢课 Docker 面板。引擎基于 HDU-KillCourse，外壳为多账号并发常驻调度引擎 + 内嵌 Web 控制面板。

## 目录结构

```
hdu-grabber/
├── docker-compose.yml        # Docker 编排配置
├── deploy/
│   ├── caddy/Caddyfile       # Caddy TLS & BasicAuth 模板
│   └── docker-compose.yml
├── data/                     # 持久化数据（挂载进容器 /app/data）
│   ├── global.json           # 全局设置（模式/定时/通知）
│   ├── accounts.json         # 账号队列（凭据+课程表+cookie）
│   ├── courses/{id}/         # 每账号课程库 course.json + 导出Excel
│   ├── clientbody/{id}.json  # 每账号选课配置缓存
│   └── log_files/            # app.log / debug.log
├── Dockerfile
├── go.mod
└── cmd/                      # 入口程序
```

## Docker Compose 部署

```bash
# 1. 克隆代码
git clone https://github.com/lizhikeer/HDU-KillCourse-LinuxDocker.git hdu-grabber
cd hdu-grabber

# 2. 配置 Caddy 反代与访问凭据（可选）
# 编辑 deploy/caddy/Caddyfile，设置你的域名与 Basic Auth 密码

# 3. 构建并启动容器
docker compose up -d --build

# 4. 验证服务状态
curl -s http://127.0.0.1:40001/api/overview
```

## 面板使用流程

1. 打开管理面板（本地回环 `http://127.0.0.1:40001` 或公网 Caddy 代理端口）；
2. 点击 **【＋ 添加账号】**，输入统一身份认证 CAS 账密或正方新教务账密；
3. 点击 **【🧪 登录】**，验证账号密码及 Session 链路是否正常；
4. 选课开放前，点击 **【🔄 更新课程库】** 同步全校教学班信息；
5. 在账号编辑弹窗中填入课程教学班名称，设置选课/退课动作及优先级；
6. 设置好开始时间与轮询间隔，**建议先开启 DryRun 干跑模式验证**，确认无误后正式挂机。

## 运维命令

```bash
# 查看实时日志
docker compose logs -f grabber

# 重启面板服务
docker compose restart grabber

# 更新代码后重新构建
docker compose up -d --build
```
