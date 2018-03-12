## Introduction
An utilities package for Golang project

## Installing
```bash
go get -u github.com/jackypanster/util
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

+ Array

```go
src := []string{"a", "a", "a"}
dst := util.Uniq(src)
```

+ Redis

+ MongoDB

+ Log
