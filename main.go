package main

import (
	"flag"
	"fmt"
	hook "github.com/robotn/gohook"
	"robotTest/fps"
)

func main() {
	flag.Parse()
	fmt.Println("--- USP连点【f1】 ---")
	fmt.Println("--- M4远距连点【f2】 ---")
	fmt.Println("--- 加特林连点【f3】 ---")
	fmt.Println("--- 自动瞄准【f4】 ---")
	fmt.Println("--- 视频笔记辅助【f5】 ---")
	fmt.Println("===============请勿用于游戏作弊，否则责任自行承担===============")

	evChan := hook.Start()
	defer hook.End()

	fpsSubject := fps.FSubject{
		Title: "fps辅助",
		L:     map[string]fps.Observer{},
	}
	// 强制开启的观察者可以提前挂载

	for ev := range evChan {
		needSend := true
		// 自己监听 注册和删除监听者
		if ev.Kind == hook.KeyHold {
			switch hook.RawcodetoKeychar(ev.Rawcode) {
			case "f1":
				needSend = false
				fpsSubject.Toggle(fps.NewUspObserver())
				break
			case "f2":
				needSend = false
				fpsSubject.Toggle(&fps.M4Observer{})
				break
			case "f3":
				needSend = false
				break
			case "f4":
				needSend = false
				break
			case "f5":
				observer := fps.VideoNoteSwitchObserver{}
				observer.Receive(ev)
				break
			}
		}
		// 其他消息全部转发给观察者
		if needSend {
			fpsSubject.Send(ev)
		}
	}

}
