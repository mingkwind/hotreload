package hotreload

import "sync"

const (
	watchDirectory = "conf/"
)

// 已注册的需要监控的文件
var callbackTable sync.Map

// CallbackFunc 配置回调函数
// filename：文件相对路径
type CallbackFunc func(filename string) error
