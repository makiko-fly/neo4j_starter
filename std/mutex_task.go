package std

import (
	"gitlab.wallstcn.com/matrix/xgbkb/std/redislogger"
)

type MutexTask struct {
	Mutex MutexInterface
	Func  func() error
}

func NewMutexTask(function func() error) *MutexTask {
	mutexTask := new(MutexTask)
	mutexTask.Func = function
	return mutexTask
}

func (self *MutexTask) Run() {
	if self.Mutex == nil {
		redislogger.Errf("MutexTask.Run(), Mutex is not set")
	}

	if success, err := self.Mutex.Lock(); err != nil {
		redislogger.Errf("MutexTask.Run(), obtain lock err: %v", err)
	} else if success {
		redislogger.Errf("MutexTask.Run(), obtain lock success!")
		defer self.Mutex.UnLock()
		self.Func()
	} else {
		redislogger.Printf("MutexTask.Run(), obtain lock fails")
	}
}

func (self *MutexTask) WithMutex(mutex MutexInterface) *MutexTask {
	self.Mutex = mutex
	return self
}
