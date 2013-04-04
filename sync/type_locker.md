#type Locker

Locker是一个接口类型。

	type Locker interface {
   		Lock()
   		Unlock()
    }
    
你可以在mutux.go中找到它的定义。

它提供了一个抽象接口，被Mutex和RWMutex实现，是Cond的成员。