package utils

type Timer struct {
	Tick  int
	Pause bool
}

func InitTimer() *Timer {
	return &Timer{
		Tick:  0,
		Pause: false,
	}
}

func (t *Timer) Update() {
	if t.Pause == false {
		t.Tick += 1
	}
}

func (t *Timer) Interval(ms int) {
	if t.Tick == ms {
		t.Pause = true
	}
}

func (t *Timer) Reset() {
	t.Tick = 0
	t.Pause = false
}
