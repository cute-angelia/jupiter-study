.PHONY: proto clean run-etcd-server

# protoc --proto_path=. --micro_out=. --go_out=. $$f; \
# jupiter protoc --out=. --file $$f -g; \
# protoc --go_out="plugins=grpc:." --proto_path=. $$f; \
# 			protoc --proto_path=. --gofast_out=plugins=grpc:. $$f; \
# dirPath=`dirname $$f`; \
# proto 生成文件
proto:
	@for d in proto; do \
        echo ===================================== library ============================================; \
		for f in $$d/**/*.proto; do \
			protoc --proto_path=. --gofast_out=plugins=grpc:. $$f; \
			echo compiled1: $$f; \
		done \
	done

jupiter:
	@for d in proto; do \
        echo ===================================== library ============================================; \
		for f in $$d/**/*.proto; do \
			jupiter protoc --out=. --file $$f -g; \
			jupiter protoc -s --out=./test/grpc --file $$f; \
			echo compiled1: $$f; \
		done \
	done

# 清理 proto 生成文件
clean:
	@for d in proto; do \
        echo ===================================== library ============================================; \
		for f in $$d/**/*.go; do \
		  	rm $$f; \
			echo clean: $$f; \
		done \
	done
	@for d in proto; do \
        echo ===================================== library ============================================; \
		for f in $$d/**/**/*.go; do \
		  	rm $$f; \
			echo clean: $$f; \
		done \
	done

up:
	git add .
	git commit -am "up"
	git pull origin master
	git push origin master
	@echo "\n 代码更新中..."

run-etcd-server:
	go run main.go --config="etcdv3://127.0.0.1:2379?key=config-study"
