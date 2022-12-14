package main

import (
	"bufio"
	"fmt"
	"os"
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
			case TIME_RUN:
				fmt.Println("计时器运行中")
				obj.now = time.Now().Unix()
				obj = cal(obj)
			case TIME_PAUSE:
				fmt.Println("计时器已暂停")
				obj = cal(obj)
			}

			fmt.Println(obj.output)
		case "g":
			switch obj.flag {
			case TIME_RUN:
				fmt.Println("计时器已启动")
			case TIME_PAUSE:
				obj.flag = TIME_RUN
				obj.start = time.Now().Unix()
				obj.now = time.Now().Unix()
				fmt.Println("恢复计时")
			case TIME_DIE:
				fmt.Println("计时开始")
				obj.flag = TIME_RUN
				obj.start = time.Now().Unix()
				obj.now = time.Now().Unix()
			}
			break
		case "e":
			switch obj.flag {
			case TIME_DIE:
				fmt.Println("计时器未启动，请先启动")
			case TIME_RUN:
				obj.now = time.Now().Unix()
			}
			fmt.Println("结束计时")
			obj.flag = TIME_DIE
			obj = cal(obj)
			goto END
		case "p":
			switch obj.flag {
			case TIME_DIE:
				fmt.Println("计时器未启动，请先启动")
			case TIME_PAUSE:
				fmt.Println("计时器已暂停")
			default:
				fmt.Println("计时器暂停")
				obj.flag = TIME_PAUSE
				obj.now = time.Now().Unix()
			}
			obj = cal(obj)
		default:
			fmt.Println("no command")
		}
		//记录到文件中
		record(&obj)
	}
END:
	fmt.Println(obj.output)
}

func cal(obj model) model {
	duration := obj.now - obj.start
	obj.hour += duration / 3600
	obj.min += (duration % 3600) / 60
	obj.sec += (duration % 3600) % 60

	if obj.sec >= 60 {
		obj.min += obj.sec / 60
		obj.sec = obj.sec % 60
	}

	if obj.min >= 60 {
		obj.hour += obj.min / 60
		obj.min = obj.min % 60
	}

	obj.output = fmt.Sprintf("持续时间 %d 小时 %d 分钟 %d 秒", obj.hour, obj.min, obj.sec)
	obj.start = time.Now().Unix()
	obj.now = time.Now().Unix()
	return obj
}

//写入文件
func record(obj *model) {
	var file *os.File
	defer file.Close()
	*obj = cal(*obj)
	if file, err := os.OpenFile("/tmp/time.log", os.O_APPEND|os.O_WRONLY, 0666); err != nil {
		panic(err)
	} else {
		write := bufio.NewWriter(file)
		_, err = write.WriteString(fmt.Sprintf("%s %s\n", time.Now().String(), obj.output))
		if err != nil {
			fmt.Println("err msg:", err)
			return
		}
		err = write.Flush()
		if err != nil {
			fmt.Println("err msg:", err)
			return
		}
	}
}
