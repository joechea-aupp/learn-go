package main

import "fmt"

type PowerDrawer interface {
	Draw(power int)
}

type socket struct{}

func (socket) Plug(device PowerDrawer, power int) {
	device.Draw(power)
}

type mobile struct {
	brand string
}

func (m mobile) Draw(power int) {
	fmt.Printf("%T -> brand: %s, power: %d\n", m, m.brand, power)
}

type laptop struct {
	cpu string
}

func (l laptop) Draw(power int) {
	fmt.Printf("%T -> cpu: %s, power: %d\n", l, l.cpu, power)
}

type toaster struct {
	amount int
}

func (t toaster) Draw(power int) {
	fmt.Printf("%T -> amount: %d, power: %d\n", t, t.amount, power)
}

type kettle struct {
	quatity string
}

func (k kettle) Draw(power int) {
	fmt.Printf("%T -> quatity: %s, power: %d\n", k, k.quatity, power)
}

func main() {
	s := socket{}
	m := mobile{"Samsung"}
	l := laptop{"Intel"}
	t := toaster{2}
	k := kettle{"1.7L"}

	s.Plug(m, 5)
	s.Plug(l, 10)
	s.Plug(t, 15)
	s.Plug(k, 20)
}
