package safemap

// 操作安全的map structure
type cmdMapData struct {
	cmdTyp  commandType                   // 操作类型
	key     string                        // 操作的key
	value   interface{}                   // 操作的value
	result  chan<- interface{}            // 操作返回的result，查找时返回结果，len返回的结果
	data    chan<- map[string]interface{} // map存储的数据，关闭channel时返回map数据
	updater UpdateFunc                    // 更新时设置更新函数
}

type UpdateFunc func(bool) interface{}

// 操作类型
type commandType int

const (
	insert commandType = iota
	del
	find
	length
	update
	closed
)

// find 结果
type findData struct {
	value interface{}
	found bool
}
