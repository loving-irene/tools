package main

import (
	"fmt"
	"time"
)

const (
	TIME_DIE   uint8 = 0
	TIME_RUN   uint8 = 1
	TIME_PAUSE uint8 = 2
)

type model struct {
	start  int64  // 开始时间
	now    int64  // 当前时间
	flag   uint8  // 标识
	hour   int64  // 小时
	min    int64  // 分钟
	sec    int64  // 秒
	output string // 输出
}

func main() {
	usage := `command comments:
s 			# check status now
g			# time start
e			# time end,print statistic
p			# time pause

print your command:
`
	fmt.Println(usage)

	var command string
	var obj model
	flag := 1

	for true {
		fmt.Scan(&command)

		switch command {
		case "s":
			if obj.flag == 0 {
				fmt.Println("计时器未启用")
			}
			fmt.Println("[状态]")
			switch obj.flag {
			case TIME_DIE:
				fmt.Println("计时器已结束/未启动")
				break
			case TIME_RUN:
				fmt.Println("计时器运行中")
				obj.now = time.Now().Unix()
				obj = cal(obj)
				break
			case TIME_PAUSE:
				fmt.Println("计时器已暂停")
				obj = cal(obj)
				break
			}

			fmt.Println(obj.output)
			break
		case "g":
			switch obj.flag {
			case TIME_RUN:
				fmt.Println("计时器已启动")
				break
			case TIME_PAUSE:
				obj.flag = TIME_RUN
				obj.start = time.Now().Unix()
				obj.now = time.Now().Unix()
				fmt.Println("恢复计时")
				break
			case TIME_DIE:
				fmt.Println("计时开始")
				obj.flag = TIME_RUN
				obj.start = time.Now().Unix()
				obj.now = time.Now().Unix()
			}
			break
		case "e":
			if obj.flag == TIME_DIE {
				fmt.Println("计时器未启动，请先启动")
				break
			}
			fmt.Println("结束计时")
			obj.flag = TIME_DIE
			obj.now = time.Now().Unix()
			flag = 2
			obj = cal(obj)
			break
		case "p":
			switch obj.flag {
			case TIME_DIE:
				fmt.Println("计时器未启动，请先启动")
				break
			case TIME_PAUSE:
				fmt.Println("计时器已暂停")
				break
			default:
				fmt.Println("暂停计时器")
			}
			obj.flag = TIME_PAUSE
			obj.now = time.Now().Unix()
			obj = cal(obj)
			break
		default:
			fmt.Println("no command")
		}

		if flag == 2 {
			break
		}

	}

	fmt.Println(obj.output)
}

func cal(obj model) model {
	duration := obj.now - obj.start
	obj.hour += duration / 3600
	obj.min += (duration % 3600) / 60
	obj.sec += (duration % 3600) % 60
	obj.output = fmt.Sprintf("持续时间 %d 小时 %d 分钟 %d 秒", obj.hour, obj.min, obj.sec)
	obj.start = time.Now().Unix()
	obj.now = time.Now().Unix()
	return obj
}
