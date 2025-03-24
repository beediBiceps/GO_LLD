package main

func main(){
    fifo:=&fifo{}
    cache:=initCache(fifo)

    cache.add("A","12")
    cache.add("b","12")
    
    lru:=&Lru{}
    cache.setEvictionAlgo(lru)
}