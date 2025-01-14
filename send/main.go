package main

import (
	"bufio"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/joho/godotenv"
	"io"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// 连接服务器
	conn, err := net.Dial("tcp", os.Getenv("address"))
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		// 提示用户输入
		fmt.Print("请输入要发送的文件路径（完成后按回车，输入 'exit' 退出）: ")

		// 读取用户输入
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}

		// 打开要传输的文件
		filePath := input
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
			continue
		}
		defer file.Close()

		// 获取文件大小
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Println("Error getting file info:", err)
			continue
		}
		fileSize := fileInfo.Size()

		// 发送文件名称和大小
		fileName := filepath.Base(filePath)
		_, err = conn.Write([]byte(fileName + "\n" + strconv.FormatInt(fileSize, 10) + "\n"))
		if err != nil {
			fmt.Println("Error sending file metadata:", err)
			continue
		}

		//初始化进度条
		bar := pb.Full.Start64(fileSize)
		//创建代理读取器
		barReader := bar.NewProxyReader(file)

		// 发送文件内容
		_, err = io.Copy(conn, barReader)
		if err != nil {
			fmt.Println("Error sending file content:", err)
			continue
		}
		//最终刷新进度条
		bar.Finish()
		fmt.Println("File sent successfully:", fileName)
	}
}
