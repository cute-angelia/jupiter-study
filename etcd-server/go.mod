module jupiter-study/etcd-server

go 1.14

require (
	github.com/douyu/jupiter v0.2.7
	jupiter-study/library v0.0.0-00010101000000-000000000000
	jupiter-study/proto v0.0.0-00010101000000-000000000000
)

replace jupiter-study/proto => ../proto

replace jupiter-study/library => ../library
