package service

var T *TaskServer

//启动服务内容
func Start() error {
	go func() {
		T = &TaskServer{}
		T.Start()
	}()
	return nil
}
