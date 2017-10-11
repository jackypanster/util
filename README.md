# Library util
An utilities package for Golang project

## Installing
```bash
go get github.com/jackypanster/util
```

## Examples

+ Queue

```go
util.Init(128, 1024)
util.JobQueue <- util.Job {
    Do: func() error {
        // do something
        return nil
    }
}
```

+ Redis

+ MongoDB