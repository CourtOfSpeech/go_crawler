package scheduler

import "crawler/engine"

//QueudeScheduler 队列 scheduler
type QueudeScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

//Submit engine.Scheduler的实现
func (s *QueudeScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

//WorkerChan engine.WorkerChan的实现
func (s *QueudeScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

//WorkerReady 接收worker 传入的channel engine.WorkerReady的实现
func (s *QueudeScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

//Run 总控 engine.Run的实现
func (s *QueudeScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			//当2种队列都有在值的时候，就将数据发送给对应的接收者
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			//2种channel并没有先后顺序，这里通过select 来判断该谁
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				//发送成功后，就从队列里面拿掉
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}

		}
	}()
}
