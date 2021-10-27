# Socks5-Proxy
![Version](https://img.shields.io/badge/Version-1.0-blue) ![License](https://img.shields.io/badge/license-MIT-yellow) ![golang](https://img.shields.io/badge/Golang-orange)

简介：Golang 实现的 Socks5 代理服务端，支持用户名、密码验证。

## 安装

### 下载
```
git clone https://github.com/truda8/Socks5-Proxy.git
```

### 构建
```bash
go build -o socks5proxy main.go
```

### 添加执行权限
```bash
chmod +x socks5proxy
```

## 使用

### 参数
 - --host=监听地址，默认：127.0.0.1
 - --port=监听端口，默认：1080
 - --user=用户名，默认：admin
 - --pass=密码，默认：admin

### 运行
```bash
./socks5proxy --host=0.0.0.0 --port=1080 --user=admin --pass=123456
```

### 后台运行
```bash
nohup ./socks5proxy > /tmp/socks5proxy.log 2>&1 &
```
