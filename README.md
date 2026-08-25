基于 Go 实现的光纤拉丝塔调度 Web 项目，一款后端服务，处理炉温拉速设定、出站指令与拉丝批次状态流转。

# FiberDraw Tower

光纤拉丝塔 / 炉温拉速指令扇出投递

Read `PROJECT.md` for the product scope, who uses it, and what the records mean. This is not a generic CMS.

## Run

```bash
set GOTOOLCHAIN=local
set CGO_ENABLED=0
go test ./... -count=1
go run ./cmd/server -addr 127.0.0.1:8080 -db ./data.sqlite
go run ./cmd/seed -db ./data.sqlite
```

Open http://127.0.0.1:8080/ for the console.
