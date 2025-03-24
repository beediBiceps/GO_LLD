package main

type weatherStation struct {
    observers []IObserver
    temperature float64
}

func (w *weatherStation) add(observer IObserver) {
    w.observers = append(w.observers, observer)
}

func (w *weatherStation) remove(observer IObserver) {
    for k, v := range w.observers {
        if v == observer {
            w.observers = append(w.observers[:k], w.observers[k+1:]...)
        }
    }
}

func (w *weatherStation) notify() {
    for _, v := range w.observers {
        v.update()
    }
}

func (w *weatherStation) getTemperature() float64 {
    return w.temperature
}
