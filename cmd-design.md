# 命令清单

## agenda help

列出命令说明

## agenda register

注册 Agenda 账号

### 语法

agenda register [flags]

- -m, --email string      邮箱
- -h, --help              帮助
- -p, --password string   密码
- -n, --phone string      电话
- -u, --username string   用户名

### 实例

agenda register -u test -p testpass -m test@test.com -n 12345678909

## agenda login

登录 Agenda 账号

### 语法

agenda login [flags]

- -h, --help              帮助
- -p, --password string   密码
- -u, --username string   用户名

### 实例

agenda login -u test -p testpass

## agenda logout

登出 Agenda 账号

### 语法

agenda logout [flags]

- -h, --help              帮助

### 实例

agenda logout
