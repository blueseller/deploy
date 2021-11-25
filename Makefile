run:
	go run cmd/root.go interactive cmd/config-dev.yml

agent-master:
	go run cmd/root.go agent-master cmd/config-dev.yml

agent-replice:
	go run cmd/root.go agent cmd/config-dev.yml -a=10.0.16.12:15151
	#go run cmd/root.go agent cmd/config-dev.yml -address=10.0.16.12:15152
