package lru

type LinkedList struct {
	key, value int
	pre, next  *LinkedList
}

type Cache struct {
	capacity   int
	caches     map[int]*LinkedList
	head, tail *LinkedList
}

func Constructor(capacity int) Cache {
	head := &LinkedList{-1, -1, nil, nil}
	tail := &LinkedList{-1, -1, nil, nil}
	head.next = tail
	tail.pre = head
	cache := Cache{capacity, make(map[int]*LinkedList), head, tail}
	return cache
}

func (this *Cache) RemoveNode(cache *LinkedList) {
	cache.pre.next = cache.next
	cache.next.pre = cache.pre
}
func (this *Cache) AddNode(node *LinkedList) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next = node
	node.next.pre = node
}
func (this *Cache) MoveToHead(cache *LinkedList) {
	this.RemoveNode(cache)
	this.AddNode(cache)
}

func (this *Cache) Get(key int) int {
	if cache, ok := this.caches[key]; ok {
		this.MoveToHead(cache)
		return cache.value
	}
	return -1
}

func (this *Cache) Put(key int, value int) {
	if cache, ok := this.caches[key]; ok {
		cache.value = value
		this.MoveToHead(cache)
		return
	}
	n := &LinkedList{key, value, nil, nil}
	if len(this.caches) >= this.capacity {
		delete(this.caches, this.tail.pre.key)
		this.RemoveNode(this.tail.pre)
	}
	this.caches[key] = n
	this.AddNode(n)
}
