package util

import "time"

const (
	HIGH = iota
	NORMAL
	LOW
)

type Task struct {
	ID       string    `json:"id"`
	Action   string    `json:"action"`
	Priority int       `json:"priority"`
	Time     time.Time `json:"time"`
	Content  string    `json:"content"`
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

func (self *TaskService) Enq(task *Task) error {
	CheckCondition(task == nil, "task should not be nil")
	return self.Rpush(task)
}

func (self *TaskService) Deq() (*Task, error) {
	reply, err := self.Lpop()
	if err != nil {
		return nil, err
	}
	if len(reply) == 0 {
		return nil, nil
	}
	var task Task
	ToInstance(reply, &task)
	return &task, nil
}
