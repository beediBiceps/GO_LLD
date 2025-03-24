package main

import "fmt"

type Customer struct
{
    id string
}

func (c *Customer) update(){
    fmt.Printf("Sending email to customer %s\n", c.id)
}
