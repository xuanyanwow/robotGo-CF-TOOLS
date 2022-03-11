package fps

import (
	"flag"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"math/rand"
)

var m4MaxShut = flag.Int("m4MaxShut", 15, "M4最大连发")
var m4MinShut = flag.Int("m4MinShut", 6, "M4最小连发")
var m4MaxSleep = flag.Int("m4MaxSleep", 80, "M4最大延迟")
var m4MinSleep = flag.Int("m4MinSleep", 120, "M4最小延迟.")

type M4Observer struct {
	can bool
}

func (o *M4Observer) Name() string {
	return "M4观察者"
}

func (o *M4Observer) Receive(ev hook.Event) {
	if ev.Kind == hook.KeyDown {
		switch hook.RawcodetoKeychar(ev.Rawcode) {
		case "1":
			o.can = true
			fmt.Println("开启M4远距")
			break
		case "2", "3":
			fmt.Println("关闭M4远距")
			o.can = false
			break
		}

	}
	if ev.Kind == hook.MouseDown && ev.Button == 2 {
		if !o.can {
			return
		}
		m4Shut()
	}
}

func m4Shut() {
RAND:
	shut := rand.Intn(*m4MaxShut)
	if shut < *m4MinShut {
		goto RAND
	}

	for i := 0; i < shut; i++ {
		robotgo.Click()
	RandSleep: // 随机暂停时间

		sleep := rand.Intn(*m4MaxSleep)
		if sleep < *m4MinSleep {
			goto RandSleep
		}
		robotgo.MicroSleep(float64(sleep))
	}
}
