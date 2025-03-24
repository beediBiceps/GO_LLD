package main

type IWeatherStation interface {
    add(observer IObserver)
    remove(observer IObserver)
    notify()
    getTemperature() float64
}