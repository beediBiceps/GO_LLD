package main

type Cache struct{
    storage map[string]string
    evictionAlgo EvictionAlgo 
    currCap int
    maxCap int
}

func initCache(e EvictionAlgo) *Cache{
    return &Cache{
        storage: make(map[string]string),
        evictionAlgo: e,
        currCap: 0,
        maxCap: 10,
    }
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo){
    c.evictionAlgo=e
}

func (c *Cache) add(key, value string){
    if c.currCap==c.maxCap{
        c.evict()
    }
    c.storage[key]=value
    c.currCap--
}

func (c *Cache) get(key string) {
    delete(c.storage, key)
}

func (c *Cache) evict(){
    c.evictionAlgo.evict(c)
    c.currCap--
}