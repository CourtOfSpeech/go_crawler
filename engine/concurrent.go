package engine

// ConcurrentEngine ConcurrentEngine
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

//Scheduler Scheduler
type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

//ReadyNotifier ReadyNotifier
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

//Run 根据传入的URL做事
func (e *ConcurrentEngine) Run(seeds ...Request) {
	//Scheduler 分发给worker
	out := make(chan ParserResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	//把所有请求都给 Scheduler
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	//接收worker返回的结果
	for {
		result := <-out
		for _, item := range result.Items {
			go func(item interface{}) {
				e.ItemChan <- item
			}(item)
		}

		//把返回的所有请求在传给 Scheduler
		for _, request := range result.Requests {
			//先去重复
			if isDuplicate(request.URL) {
				continue
			}
			//没有重复在提交给Scheduler
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			//tell scheduler I'm  ready
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var urlVisitRecord = make(map[string]bool)

//去重复判断
func isDuplicate(url string) bool {
	if urlVisitRecord[url] {
		return true
	}
	urlVisitRecord[url] = true
	return false
}
