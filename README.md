## Introduction
An utilities package for Golang project

## Installing
```bash
go get github.com/jackypanster/util
```

## Examples

+ Queue

```go
util.InitQueue(64, 65536)
util.JobQueue <- util.Job {
    Do: func() error {
        // do something
        return nil
    }
}
```

+ Redis

+ MongoDB