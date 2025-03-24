package main

import (
    "fmt"
)

type Item struct{
    observerList []IObserver
    Name string
    inStock bool
}

func (i *Item) add(o IObserver){
    i.observerList=append(i.observerList, o)
}

func (i *Item) remove(o IObserver){
    for k,v:=range i.observerList{
        if v==o{
            i.observerList=append(i.observerList[:k], i.observerList[k+1:]...)
        }
    }
}

func (i *Item) notify(){
    for _,v:=range i.observerList{
        v.update()
    }
}

