package object

import (
	"fmt"
	"sync"
)

import log "github.com/Deansquirrel/goToolLog"

type ITaskManager interface {
	GetIdList() []string
	GetTask(id string) ITask
	GetChRegister() chan<- ITask
	GetChUnregister() chan<- string
	GetChClose() chan<- struct{}
}

type taskManager struct {
	task         map[string]ITask
	chRegister   chan ITask
	chUnregister chan string
	chClose      chan struct{}

	lock sync.Mutex
}

func NewTaskManager() ITaskManager {
	t := taskManager{
		task:         make(map[string]ITask),
		chRegister:   make(chan ITask),
		chUnregister: make(chan string),
		chClose:      make(chan struct{}),
	}
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

func (t *taskManager) GetChClose() chan<- struct{} {
	return t.chClose
}

func (t *taskManager) start() {
	go func() {
		for {
			select {
			case task := <-t.chRegister:
				t.register(task)
			case task := <-t.chUnregister:
				t.unregister(task)
			case <-t.chClose:
				list := t.GetIdList()
				for _, id := range list {
					t.GetChUnregister() <- id
				}
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
