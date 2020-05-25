package baLogger

import "time"

type Decorator struct {
	types []string
}

func (d *Decorator) Flag(flag string) *Decorator {

	return d
}

func (d *Decorator) Bench(timeVal time.Time) *Decorator {
	return d
}

func Log(f func()) Decorator {
	decorator := Decorator{}
	return decorator
}
