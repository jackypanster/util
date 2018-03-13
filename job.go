package util

import (
	"time"

	"github.com/satori/go.uuid"
)

const (
	HIGH uint = iota
	NORMAL
	LOW
)

type Task struct {
	ID       string      `json:"id"`
	Priority uint        `json:"priority"`
	Time     time.Time   `json:"time"`
	Retries  uint        `json:"retries"`
	Content  interface{} `json:"content"`
}

func NewTask(content interface{}) Task {
	id, err := uuid.NewV4()
	CheckErrf(err, "unable to gen UUID")
	return Task{
		ID:       id.String(),
		Time:     time.Now(),
		Retries:  0,
		Content:  content,
		Priority: NORMAL,
	}
}

type TaskService struct {
	*RedisService
}

func NewTaskService(redisService *RedisService) *TaskService {
	CheckCondition(redisService == nil, "redisService should not be nil")
	return &TaskService{
		redisService,
	}
}

func (self *TaskService) Enq(content interface{}) error {
	CheckCondition(content == nil, "content should not be nil")
	return self.Rpush(NewTask(content))
}

func (self *TaskService) Deq() (interface{}, error) {
	reply, err := self.Lpop()
	if err != nil {
		return nil, err
	}
	if len(reply) == 0 {
		return nil, nil
	}
	var task Task
	ToInstance(reply, &task)
	return task.Content, nil
}
