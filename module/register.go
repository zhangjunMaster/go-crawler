package module

/*
 * 根据基础接口定义的接口
 * 接口定义的方法，返回的数据还是接口类型的数据
 */
type Register interface {
	// 注册组件实例
	Register(module Module) (bool, error)
	// 注销主键实例
	Unregister(mid MID) (bool, error)
	// 获取一个该组件类型的组件实例
	Get(moduleType Type) (Module, error)
	// 获取指定类型的所有组件
	GetAllByType(moduleType Type) (map[MID]Module, error)
	// 获取所有组件实例
	GetAll() (map[MID]Module, error)
	// 清除所欲的组件注册记录
	Clear()
}
