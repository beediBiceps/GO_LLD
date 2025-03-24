package main

type Duck struct {
	quack Iquack
	fly   IFly
}

func (d *Duck) implementFly() {
	d.fly.fly()
}

func (d *Duck) implementQuack() {
	d.quack.quack()
}

func (d *Duck) setFly(fly IFly) {
	d.fly = fly
}

func (d *Duck) setQuack(quack Iquack) {
	d.quack = quack
}
