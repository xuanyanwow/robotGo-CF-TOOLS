package fps

import hook "github.com/robotn/gohook"

// Observer 抽象观察者
type Observer interface {
	Receive(ev hook.Event)
	UnObserver()
	Name() string
}
