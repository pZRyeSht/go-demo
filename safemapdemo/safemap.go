package safemap

/**
并发安全的map
**/

// SafeMapImp 定义一个并发安全的map，操作interface
type SafeMapImp interface {
	Insert(string, interface{})
	Delete(string)
	Find(string) (interface{}, bool)
	Len() int
	Update(string, UpdateFunc)
	Close() map[string]interface{}
}

// 基于cmdMapData类型的channel实现一个可发送、可接收的并发安全的map
type safeMap chan cmdMapData

func (s2 safeMap) Insert(key string, value interface{}) {
	s2 <- cmdMapData{
		cmdTyp: insert,
		key:    key,
		value:  value,
	}
}

func (s2 safeMap) Delete(key string) {
	s2 <- cmdMapData{
		cmdTyp: del,
		key:    key,
	}
}

func (s2 safeMap) Find(key string) (interface{}, bool) {
	res := make(chan interface{})
	s2 <- cmdMapData{
		cmdTyp: find,
		key:    key,
		result: res,
	}
	resp := (<-res).(findData)
	return resp.value, resp.found
}

func (s2 safeMap) Len() int {
	res := make(chan interface{})
	s2 <- cmdMapData{
		cmdTyp: length,
		result: res,
	}
	return (<-res).(int)
}

func (s2 safeMap) Update(key string, updater UpdateFunc) {
	s2 <- cmdMapData{
		cmdTyp:  update,
		key:     key,
		updater: updater,
	}
}

func (s2 safeMap) Close() map[string]interface{} {
	res := make(chan map[string]interface{})
	s2 <- cmdMapData{
		cmdTyp: closed,
		data:   res,
	}
	return <-res
}

func (s2 safeMap) run() {
	mp := make(map[string]interface{})
	for cmd := range s2 {
		switch cmd.cmdTyp {
		case insert:
			mp[cmd.key] = cmd.value
		case del:
			delete(mp, cmd.key)
		case find:
			value, ok := mp[cmd.key]
			cmd.result <- findData{
				value: value,
				found: ok,
			}
		case length:
			cmd.result <- len(mp)
		case update:
			_, ok := mp[cmd.key]
			mp[cmd.key] = cmd.updater(ok) // updater 不用调用safeMap的其他方法，会造成死锁
		case closed:
			close(s2)
			cmd.data <- mp
		}
	}
}

// NewSafeMap safeMap的工厂方法
func NewSafeMap() SafeMapImp {
	smp := make(safeMap)
	go smp.run()
	return smp
}
