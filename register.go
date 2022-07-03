package hotreload

// Register 注册监控对象
func Register(filename string, fn CallbackFunc) error {
	_, ok := callbackTable.Load(filename)
	if ok {
		panic("filename already exists.")
	}
	callbackTable.Store(filename, fn)

	return nil
}
