package main

import "fmt"

type fifo struct {
}

func (f *fifo) evict(c *Cache) {
	   fmt.Println("Evicting by fifo strtegy")
}