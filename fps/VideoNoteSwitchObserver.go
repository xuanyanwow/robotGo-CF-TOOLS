package fps

import (
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

type VideoNoteSwitchObserver struct {
	can bool
}
func (o *VideoNoteSwitchObserver) Name() string {
	return "视频笔记切换观察者"
}

func (o *VideoNoteSwitchObserver) Receive(ev hook.Event) {
	//在这里监听按键 shut
	if ev.Kind == hook.KeyHold {
		switch hook.RawcodetoKeychar(ev.Rawcode) {
		case "f5":
			// 记录当前位置
			x,y := robotgo.GetMousePos()

			// 切换到左边屏幕  点击
			robotgo.Move(300,300)
			robotgo.Click()
			robotgo.MicroSleep(100)

			// 回到原来位置 点击
			robotgo.Move(x,y)
			robotgo.Click()
			break
		}
	}
}