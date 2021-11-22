package chief

import (
	"context"
	"fmt"
	"log"
	"net"

	commandPb "github.com/blueseller/deploy.git/api/agent/command/v1"
	"google.golang.org/grpc"
)

func Run(ctx context.Context) error {
	port := "" // todo use ctx
	ip := ""   // todo use ctx

	// 获取本机ip地址
	if ip != "" {
		ip = GetIP()
	}

	if ip == "" {
		return fmt.Errorf("获取本机IP地址失败,请使用参数传入一个IP地址 ")
	}

	server := grpc.NewServer()
	commandPb.RegisterStreamCommandSerivceServer(server, &CommandServices{})

	lis, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)

	}
	server.Serve(lis)
	return nil
}

func GetIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, address := range addrs {
		// 检查 ip 地址判断是否回环地址
		if ipnet, flag := address.(*net.IPNet); flag && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
