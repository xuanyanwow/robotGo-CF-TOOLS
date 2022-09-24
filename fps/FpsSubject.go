package fps

import (
	"fmt"
	hook "github.com/robotn/gohook"
)

type FSubject struct {
	Title string
	L     map[string]Observer
}

func (sub *FSubject) Add(o Observer) {
	_, ok := sub.L[o.Name()]
	if !ok {
		sub.L[o.Name()] = o
	}
}

func (sub *FSubject) Remove(o Observer) {
	fmt.Println(o)
	_, ok := sub.L[o.Name()]
	if ok {
		delete(sub.L, o.Name())
	}
}

func (sub *FSubject) Toggle(o Observer) {
	realO, ok := sub.L[o.Name()]
	if ok {
		realO.UnObserver()
		delete(sub.L, realO.Name())
	} else {
		sub.L[o.Name()] = o
	}
}

func (sub *FSubject) Send(ev hook.Event) {

	for observerIndex := range sub.L {
		sub.L[observerIndex].Receive(ev)
	}
}
