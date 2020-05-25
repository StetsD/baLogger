package baLogger

import (
	"fmt"
	"time"
)

type Decorator struct {
	types []string
	f     func()
	flag  string
}

func (d *Decorator) Call() {
	d.f()
}

func (d *Decorator) Flag(flag string) *Decorator {
	prevFunc := d.f
	d.flag = flag
	d.f = func() {
		fmt.Printf("START : %s\n", flag)
		prevFunc()
		fmt.Printf("END : %s\n", flag)
	}
	return d
}

func (d *Decorator) Bench() *Decorator {
	prevFunc := d.f

	d.f = func() {
		start := time.Now()
		prevFunc()
		fmt.Printf("Timer : %s %s\n", d.flag, time.Since(start))
	}

	return d
}

func Log(f func()) *Decorator {
	decorator := &Decorator{make([]string, 0), f, ""}
	return decorator
}
