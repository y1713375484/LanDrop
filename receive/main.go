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
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	//保存文件的路径不等于空
	if os.Getenv("savePath") != "" {
		_, err = os.Stat(os.Getenv("savePath"))
		//检测要保存的路径是否存在
		if os.IsNotExist(err) {
			err := os.Mkdir(os.Getenv("savePath"), 0777)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	// 监听端口
	listen, err := net.Listen("tcp", os.Getenv("address"))
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listen.Close()
	fmt.Println("Server is listening on " + os.Getenv("address"))

	for {
		// 接受客户端连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn) // 使用 goroutine 处理连接
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		// 读取文件名称
		fileName, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading file name:", err)
			return
		}
		fileName = fileName[:len(fileName)-1]

		//读取文件大小
		size, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Error reading file size:", err)
				break
			}
			fmt.Println("Error reading file size:", err)
			return
		}
		size = size[:len(size)-1]
		sizeInt, _ := strconv.ParseInt(size, 10, 64)

		// 创建文件
		file, err := os.Create(filepath.Join(os.Getenv("savePath"), fileName))
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		//初始化进度条
		bar := pb.Full.Start64(sizeInt)
		//创建代理读取器
		barReader := bar.NewProxyReader(conn)
		readSum := int64(0)
		buffer := make([]byte, 1024)
		for readSum < sizeInt {
			n, err := barReader.Read(buffer)
			if err != nil {
				if err == io.EOF {
					break
				}
			}
			readSum += int64(n)
			file.Write(buffer[:n])
		}
		//最终刷新进度条
		bar.Finish()
		fmt.Println("File received successfully:", fileName)
	}
}
