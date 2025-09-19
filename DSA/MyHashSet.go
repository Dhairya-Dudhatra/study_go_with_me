type MyHashSet struct {
    keyval map[int]struct{}
}


func Constructor() MyHashSet {
    myhash := MyHashSet{}
    myhash.keyval = make(map[int]struct{}, 1000)
    return myhash
}


func (this *MyHashSet) Add(key int)  {
    this.keyval[key] = struct{}{}
}


func (this *MyHashSet) Remove(key int)  {
    if _, found := this.keyval[key]; found {
        delete(this.keyval, key)
    }  
}


func (this *MyHashSet) Contains(key int) bool {
    if _, found := this.keyval[key]; !found {
        return false
    } 
    return true
}
