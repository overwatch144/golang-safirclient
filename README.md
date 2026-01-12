# Golang Safir Client

A unified Go client library for Safir OpenStack services:
- Safir Optimization (safiroptimization)
- Safir Migration (safirmigration)
- Safir Cloud Watcher (safir_cloud_watcher)

## Installation
```bash
go get github.com/overwatch144/golang-safirclient
```

## Usage

### Safir Optimization
```go
import "github.com/overwatch144/golang-safirclient/optimization"

client, err := optimization.NewClient(optimization.ClientOptions{
    AuthURL:     "http://10.13.0.10:5000/v3",
    Username:    "admin",
    Password:    "password",
    ProjectName: "admin",
})
```

### Safir Migration
```go
import "github.com/overwatch144/golang-safirclient/migration"

client, err := migration.NewClient(migration.ClientOptions{
    AuthURL:     "http://10.13.0.10:5000/v3",
    Username:    "admin",
    Password:    "password",
    ProjectName: "admin",
})
```

### Safir Cloud Watcher
```go
import "github.com/overwatch144/golang-safirclient/cloudwatcher"

client, err := cloudwatcher.NewClient(cloudwatcher.ClientOptions{
    AuthURL:     "http://10.13.0.10:5000/v3",
    Username:    "admin",
    Password:    "password",
    ProjectName: "admin",
})
```

## License

MIT
