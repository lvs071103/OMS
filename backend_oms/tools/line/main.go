package main

import (
	"fmt"
	"os/exec"
	"sync"
)

// pingIP 尝试 ping 一个 IP 地址，并将结果发送到结果通道
func pingIP(ip string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	cmd := exec.Command("ping", "-c", "1", "-W", "1", ip)
	err := cmd.Run()
	if err != nil {
		results <- fmt.Sprintf("%s: DOWN", ip)
	} else {
		results <- fmt.Sprintf("%s: UP", ip)
	}
}

func main() {
	// 示例 IP 地址列表
	ips := []string{
		"192.168.1.1",
		"8.8.8.8",
		"8.8.4.4",
		// 添加更多 IP 地址
	}

	// 创建一个 WaitGroup 以等待所有 goroutine 完成
	var wg sync.WaitGroup
	// 创建一个通道以接收结果
	results := make(chan string, len(ips))

	// 并发 ping IP 地址
	for _, ip := range ips {
		wg.Add(1)
		go pingIP(ip, &wg, results)
	}

	// 等待所有 goroutine 完成
	wg.Wait()
	close(results)

	// 打印结果
	for result := range results {
		fmt.Println(result)
	}
}
