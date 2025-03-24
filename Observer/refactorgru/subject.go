package main

type ISubject interface {
    add(ovserver IObserver)
    remove(observer IObserver)
    notify()
}