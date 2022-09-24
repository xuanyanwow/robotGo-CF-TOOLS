package fps

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"robotTest/utils"
)

type JatelingObserver struct {
	run bool
}

func (o *JatelingObserver) Name() string {
	return "加特林观察者"
}

func (o *JatelingObserver) Receive(ev hook.Event) {
	// 只有第一次点击左键会触发
	if ev.Kind == hook.MouseDown && ev.Button == 1 && ev.Clicks == 1 {
		if !o.run {
			o.run = true
			fmt.Println("开启加特林")
		}
		go o.Shut()
	}
}

func (o *JatelingObserver) UnObserver() {
	o.run = false
}

// MouseHold -> MouseUp
// 11:24:44.1711 -> 11:24:44.2216  (2216-1711=505)
// MouseHold -> Next MouseHold
// 11:24:44.1711 -> 11:24:44.2719    (2719-2216=503)

func (o *JatelingObserver) Shut() {
	for {
		if !o.run {
			break
		}

		sleep := utils.RandBetween(200, 400)

		robotgo.Toggle("left")
		robotgo.MilliSleep(sleep)
		robotgo.Toggle("left", "up")

		sleep = utils.RandBetween(200, 400)
		robotgo.MilliSleep(sleep)
	}

	fmt.Println("关闭加特林")
}
