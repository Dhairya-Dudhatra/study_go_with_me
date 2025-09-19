type MyHashMap struct {
    keyval map[int]int
}


func Constructor() MyHashMap {
    return MyHashMap{
        keyval: make(map[int]int, 10000),
    }
}


func (this *MyHashMap) Put(key int, value int)  {
    this.keyval[key] = value
}


func (this *MyHashMap) Get(key int) int {
    if val, found := this.keyval[key]; !found {
        return -1
    } else {
        return val
    }
}


func (this *MyHashMap) Remove(key int)  {
    delete(this.keyval, key)
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */