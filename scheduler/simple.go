package scheduler

import "crawler/engine"

//SimpleScheduler SimpleScheduler
type SimpleScheduler struct {
	workerChan chan engine.Request
}

//Submit 实现 engine.Scheduler.Submit()
func (s *SimpleScheduler) Submit(r engine.Request) {
	//将请求信息，发送到worker channel,这样直接发送，容易卡死，改为goroutine
	//s.workerChan <- r
	go func() {
		s.workerChan <- r
	}()
}

//ConfigureMasterWorkerChan 实现 engine.Scheduler.ConfigureMasterWorkerChan()
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
