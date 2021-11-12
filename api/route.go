package api

import (
	"github.com/gorilla/mux"

	"github.com/blueseller/deploy/deploy/paas"
)

var r *mux.Router

func Router() {
	r = mux.NewRouter()
	// paas层部署
	paasDeploy := r.PathPrefix("/deploy/paas").Subrouter()
	paasDeploy.HandleFunc("/db/mongo", paas.DeployMongoHandler)
	paasDeploy.HandleFunc("/db/pg", paas.DeployPgHandler)
	paasDeploy.HandleFunc("/db/mysql", paas.DeployMysqlHandler)

	// paas层部署, 申请DB层资源
	// 可以从各个云平台获取一个地址
	paasDeploy.HandleFunc("/db/apply/mongo", paas.ApplyMysqlHandler)
	paasDeploy.HandleFunc("/db/apply/pg", paas.ApplyMysqlHandler)
	paasDeploy.HandleFunc("/db/apply/mysql", paas.ApplyMysqlHandler)

	// paas层部署,直接录入DB层资源
	// 可以直接把已存在的资源放入进来
	paasDeploy.HandleFunc("/db/input/mongo", paas.ApplyMysqlHandler)
	paasDeploy.HandleFunc("/db/input/pg", paas.ApplyMysqlHandler)
	paasDeploy.HandleFunc("/db/input/mysql", paas.ApplyMysqlHandler)

	// saas层部署
	saasDeploy := r.PathPrefix("/deploy/saas").Subrouter()
	saasDeploy := r.PathPrefix("/deploy/server/one").Subrouter()
	saasDeploy := r.PathPrefix("/deploy/server/only").Subrouter()
	saasDeploy := r.PathPrefix("/deploy/server/migration/one").Subrouter()
	saasDeploy := r.PathPrefix("/deploy/server/migration/all").Subrouter()
	saasDeploy := r.PathPrefix("/deploy/migration/one").Subrouter()
	saasDeploy := r.PathPrefix("/deploy/migration/one/action").Subrouter()

	// 部署进度管理
	saasDeploy := r.PathPrefix("/deploy/progress").Subrouter()

	// 配置管理
	configDeploy := r.PathPrefix("/deploy/config").Subrouter()
	configDeploy.HandleFunc("/reflesh/byfile").Subrouter()
	configDeploy.HandleFunc("/reflesh/prefix").Subrouter()
	configDeploy.HandleFunc("/diffs").Subrouter()
	configDeploy.HandleFunc("/put").Subrouter()
}
