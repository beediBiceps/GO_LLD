package main

import (
	"fmt"
)

type MobileObserver struct {
	subject     IWeatherStation
	temperature float64
}

func (m *MobileObserver) update() {
	m.temperature = m.subject.getTemperature()
	fmt.Println("Mobile: The temperature is", m.temperature)
}
