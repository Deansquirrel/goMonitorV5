package object

import (
	"fmt"
	"github.com/Deansquirrel/goMonitorV5/global"
	"sync"
)

import log "github.com/Deansquirrel/goToolLog"

type taskManager struct {
	task         map[string]ITask
	chRegister   chan ITask
	chUnregister chan string

	lock sync.Mutex
}

func NewTaskManager() *taskManager {
	t := taskManager{}
	t.start()
	return &t
}

func (t *taskManager) GetIdList() []string {
	list := make([]string, 0)
	for key := range t.task {
		list = append(list, key)
	}
	return list
}

func (t *taskManager) GetTask(id string) ITask {
	task, ok := t.task[id]
	if ok {
		return task
	}
	return nil
}

func (t *taskManager) GetChRegister() chan<- ITask {
	return t.chRegister
}

func (t *taskManager) GetChUnregister() chan<- string {
	return t.chUnregister
}

func (t *taskManager) start() {
	go func() {
		for {
			select {
			case task := <-t.chRegister:
				t.register(task)
			case task := <-t.chUnregister:
				t.unregister(task)
			case <-global.Ctx.Done():
				return
			}
		}
	}()
}

func (t *taskManager) register(task ITask) {
	t.lock.Lock()
	defer t.lock.Unlock()
	_, ok := t.task[task.GetTaskId()]
	if ok {
		log.Debug(fmt.Sprintf("task %s is already exist", task.GetTaskId()))
		return
	}
	t.task[task.GetTaskId()] = task
	log.Info(fmt.Sprintf("task %s is added", task.GetTaskId()))
	return
}

func (t *taskManager) unregister(id string) {
	t.lock.Lock()
	defer t.lock.Unlock()
	task, ok := t.task[id]
	if !ok {
		log.Debug(fmt.Sprintf("task %s is not exist", id))
		return
	}
	task.Stop()
	delete(t.task, id)
}
