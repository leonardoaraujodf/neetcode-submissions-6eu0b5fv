type LRUCache struct {
    mapKeyToElem map[int]*list.Element
	mapElemToKey map[*list.Element]int
	l *list.List
	capacity int
	size int
}

func Constructor(capacity int) LRUCache {
    return LRUCache {
		mapKeyToElem: make(map[int]*list.Element),
		mapElemToKey: make(map[*list.Element]int),
		l: list.New(),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
    e, ok := this.mapKeyToElem[key]
	if !ok {
		return -1
	}

	this.l.MoveToFront(e)
	return e.Value.(int)
}

func (this *LRUCache) Put(key int, value int) {
    e, ok := this.mapKeyToElem[key]
	if !ok {
		this.l.PushFront(value)
		this.mapKeyToElem[key] = this.l.Front()
		this.mapElemToKey[this.l.Front()] = key
		this.size++
	} else {
		e.Value = value
		this.l.MoveToFront(e)
	}

	if this.size > this.capacity {
		el := this.l.Back()
		k := this.mapElemToKey[el]
		delete(this.mapElemToKey, el)
		delete(this.mapKeyToElem, k)
		this.l.Remove(el)
		this.size--
	}
}
