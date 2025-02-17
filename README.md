# 局域网传输工具

这是一个简单的局域网文件传输工具，支持设置保存路径、查看传输进度，并通过配置文件设置连接IP地址。该工具旨在帮助用户在局域网内快速、方便地传输文件。

## 功能特性

- **设置保存路径**：用户可以自定义文件保存的路径。
- **传输进度查看**：实时显示文件传输的进度。
- **配置文件设置**：通过配置文件设置连接的IP地址，方便管理和使用。

## 使用说明

### 1. 接收端配置文件

在receive目录下找到.env文件，编辑以下内容：

```ini
address=127.0.0.1:8888 #连接监听的地址
savePath=              #保存文件的路径 
```
### 2. 发送端配置文件
在send目录下找到.env文件，编辑以下内容
```ini
address=127.0.0.1:8888 #连接监听的地址
```

### 3. 运行程序
#### 编译运行
确保当前发送端与接收端在同一局域网，可以先用ping命令进行测试
然后分别在发送和接收端目录下运行
```
go run main.go
```
#### 二进制文件直接运行
windwos直接运行.exe结尾文件即可


### 4. 运行效果
#### 发送端
```
xxx@Mac send % go run main.go
请输入要发送的文件路径（完成后按回车，输入 'exit' 退出）: /Users/Desktop/schoolvideo.zip
550.88 MiB / 550.88 MiB [--------------------------------------------------------------------------] 100.00% 483.13 MiB p/s 1.3s
File sent successfully: schoolvideo.zip
请输入要发送的文件路径（完成后按回车，输入 'exit' 退出）: ^Csignal: interrupt
```

#### 接收端
```
xxx@Mac receive % go run main.go
Server is listening on 127.0.0.1:8888
550.88 MiB / 550.88 MiB [--------------------------------------------------------------------------] 100.00% 481.34 MiB p/s 1.3s
File received successfully: schoolvideo.zip
```
