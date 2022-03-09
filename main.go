package main

import (
	"flag"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"math/rand"
)

var maxShut = flag.Int("maxShut", 4, "最大连发")
var minShut = flag.Int("minShut", 3, "最小连发")
var maxSleep = flag.Int("maxSleep", 80, "最大延迟")
var minSleep = flag.Int("minSleep", 120, "最小延迟.")

func main() {
	flag.Parse()

	fmt.Println("--- 当切换为副武器(按下2时) 开启功能 ---\n--- 切换主武器/近战武器(按下1/3时) 关闭功能 ---")

	// TODO 打印参数

	can := false
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		if ev.Kind == hook.KeyDown {
			switch hook.RawcodetoKeychar(ev.Rawcode) {
			case "1","3":
				can = false
				fmt.Println("关闭连发")
				break
			case "2":
				fmt.Println("开启连发")
				can = true
				break
			}
		}
		if ev.Kind == hook.MouseDown && ev.Button == 2 {
			if !can {
				return
			}
			shut()
		}
	}

}

func shut() {
RAND:
	shut := rand.Intn(*maxShut)
	if shut < *minShut {
		goto RAND
	}

	for i := 0; i < shut; i++ {
		robotgo.Click()
	RandSleep: // 随机暂停时间

		sleep := rand.Intn(*maxSleep)
		if sleep < *minSleep {
			goto RandSleep
		}
		robotgo.MicroSleep(float64(sleep))
	}
}
