package api

type ServerOneReq struct {
	ServiceName string `json: "service_name"`
}

type ServerOnlyReq struct {
	ServiceName string `json: "service_name"`
}

type ServerMigrationAllReq struct {
	ServiceName string `json: "service_name"`
}

type MigrationOneReq struct {
}

type MigrationOneActionReq struct {
}
