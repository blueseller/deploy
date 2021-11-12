module github.com/blueseller/deploy.git

go 1.15

require (
	github.com/blueseller/deploy v0.0.0-00010101000000-000000000000
	github.com/google/godepq v0.0.0-20190501212251-2c635fd1e5fe // indirect
	github.com/gorilla/mux v1.8.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/blueseller/deploy => ./
