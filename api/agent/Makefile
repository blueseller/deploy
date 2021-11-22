build: clean
	protoc -I/home/liukai02/go/src/ --go_out=plugins=grpc:/home/liukai02/go/src/ github.com/blueseller/deploy.git/api/agent/types/types.proto
	protoc -I/home/liukai02/go/src/ --go_out=plugins=grpc:/home/liukai02/go/src/ github.com/blueseller/deploy.git/api/agent/command/v1/command.proto
	protoc -I/home/liukai02/go/src/ --go_out=plugins=grpc:/home/liukai02/go/src/ github.com/blueseller/deploy.git/api/agent/heartbeat/v1/heartbeat.proto
	protoc -I/home/liukai02/go/src/ --go_out=plugins=grpc:/home/liukai02/go/src/ github.com/blueseller/deploy.git/api/agent/monitor/v1/monitor.proto

clean:
	find . -name '*.pb.go' -type f -print -exec rm -rf {} \;
