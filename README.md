# Hotnews

自动采集一些热点的新闻，并转换成RSS的服务内容

## 添加到系统服务(Linux/Ubuntu)

### 1. 修改文件 hotnews.service

1. 修改 ${User} 为你希望运行的用户
2. 修改 ${Group} 为你希望运行的用户组
3. 修改 ${PATH} 为你程序所在的目录
```
User=${User}
Group=${Group}
WorkingDirectory=${PATH}
```

### 2. 拷贝文件到 /etc/systemd/system 目录

```shell
sudo cp hotnews.service /etc/systemd/system
```

### 3. 启动服务

开机启动
```shell
sudo systemctl enable hotnews
```

启动服务
```shell
sudo systemctl start hotnews
```

