module github.com/blueseller/deploy.git

go 1.15

require (
	github.com/google/godepq v0.0.0-20190501212251-2c635fd1e5fe // indirect
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.8.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/blueseller/deploy.git => ./
