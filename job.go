package util

const (
	HIGH uint = iota
	NORMAL
	LOW
)

type Task struct {
	ID       string      `json:"id"`
	Priority uint        `json:"priority"`
	Time     int64       `json:"time"`
	Retries  uint        `json:"retries"`
	Content  interface{} `json:"content"`
}

func NewTask(id string, content interface{}, time int64) Task {
	return Task{
		ID:       id,
		Time:     time,
		Retries:  0,
		Content:  content,
		Priority: NORMAL,
	}
}

type TaskService struct {
	*RedisService
}

func NewTaskService(redisService *RedisService) *TaskService {
	CheckNil(redisService, "redis service should not be nil")
	return &TaskService{
		redisService,
	}
}

func (self *TaskService) Enq(id string, content interface{}, time int64) error {
	CheckStr(id, "id")
	CheckCondition(time < 0, "time should not be negative")
	CheckNil(content, "content should not be nil")
	v, err := ToJsonString(NewTask(id, content, time))
	if err != nil {
		return err
	}
	return self.Rpush(v)
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
	err = ToInstance(reply, &task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}
