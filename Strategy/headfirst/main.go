package main

func main() {
	yellowDuck := &Duck{
		fly:   Nofly{},  
		quack: Yesquack{}, 
	}

	yellowDuck.implementQuack()
	yellowDuck.implementFly()

	yellowDuck.setFly(Yesfly{})
	yellowDuck.setQuack(Yesquack{})

	yellowDuck.implementQuack()
	yellowDuck.implementFly()
}
