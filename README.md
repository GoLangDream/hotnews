# Hotnews

自动采集一些热点的新闻，并转换成RSS的服务内容

## 添加到系统服务(Linux/Ubuntu)

### 1. 修改文件 hotnews.service 

ExecStart 所在的行
1. USER 改成启动程序的用户
2. PROJECT_PATH 改成程序所在的目录

例如
```
ExecStart=runuser -l user -c 'cd /home/user/golang/hotnews && /usr/bin/sh start.sh'
```

### 2. 拷贝文件到 /etc/systemd/system 目录

```shell
cp hotnews.service /etc/systemd/system
```

### 3. 启动服务

```shell
systemctl enable hotnews --now
```

