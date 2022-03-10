package fps

import (
	"fmt"
	hook "github.com/robotn/gohook"
)

type M4Observer struct {
}

func (o *M4Observer) Name() string {
	return "M4观察者"
}

func (o *M4Observer) Receive(ev hook.Event) {
	fmt.Println("A类观察者【" + o.Name() + "】触发")
}
