# Agenda

> 课程《服务计算》作业三：用 Go 实现命令行 Agenda

## 安装

```
go get github.com/mensu/Agenda-Go
```

## config 文件

默认使用 ``$HOME/.agenda-go.yaml``

如果找不到的话，将使用如下的默认设置

```yaml
# 工作目录。其他配置如果使用相对路径，则相对该工作目录
cwd: .
# log 的路径。如果环境变量设定了 DEBUG，则 log 会输出到 stderr
log: /dev/null
# 用户数据的路径。JSON 格式
user_data: data/user_data

```
