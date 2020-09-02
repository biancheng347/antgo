package plugins

// 定义一个接口，里面有两个方法
type pluginFunc interface {
	Before() interface{} //在什么之前执行
	After()              //在什么之后执行
}

// 定义一个类，来存放我们的插件
type New struct {
	List map[string]pluginFunc
}

// 初始化插件
func (p *New) Init() {
	p.List = make(map[string]pluginFunc)
}

// 注册插件
func (p *New) Register(path, name string, plugin pluginFunc) {
	if path == name {
		p.List[name] = plugin
	}
}
