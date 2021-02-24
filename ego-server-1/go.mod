module jupiter-study/ego-server-1

go 1.14

require (
	github.com/frankban/quicktest v1.11.3 // indirect
	github.com/gotomicro/ego v0.3.8
	github.com/gotomicro/ego-component/eetcd v0.1.4
	google.golang.org/grpc v1.31.0
)

replace google.golang.org/grpc v1.31.0 => google.golang.org/grpc v1.29.1
