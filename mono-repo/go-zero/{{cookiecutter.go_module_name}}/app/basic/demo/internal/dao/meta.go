package dao

import (
	"{{cookiecutter.go_module_name}}/app/basic/demo/internal/dao/async"
	"{{cookiecutter.go_module_name}}/app/basic/demo/internal/dao/cache"
	"{{cookiecutter.go_module_name}}/app/basic/demo/internal/dao/db"
	"{{cookiecutter.go_module_name}}/app/basic/demo/internal/dao/graphql"
	"{{cookiecutter.go_module_name}}/app/basic/demo/internal/dao/http"
	"{{cookiecutter.go_module_name}}/app/basic/demo/internal/dao/mq"
	"{{cookiecutter.go_module_name}}/app/basic/demo/internal/dao/rpc"
	"{{cookiecutter.go_module_name}}/app/basic/demo/proto/config"
)

/*
数据层资源收敛入口:
	- 收敛所有数据资源对象
	- dao 层: 收敛对数据资源的处理逻辑:
		- async: all go async task
		- db: crud 操作
		- cache: get/set
		- mq: pub/sub
		- http: call external http api
		- rpc: call other biz logic
	- 不包含业务逻辑, 业务逻辑在上层 domain 内组织
	- 上帝归上帝, 凯撒归凯撒
*/
type MetaResource struct {
	Async *async.Dao   // async task handler
	DB    *db.Dao      // db layer
	Cache *cache.Dao   // cache layer
	MQ    *mq.Dao      // mq layer
	Graph *graphql.Dao // graphql
	HTTP  *http.Dao    // required http client layer
	RPC   *rpc.Dao     // required rpc client layer
}

func NewMetaResource(cfg config.Config, isRpcServer bool) *MetaResource {
	return &MetaResource{
		Async: async.New(),
		DB:    db.New(cfg),
		Cache: cache.New(cfg.Cache),
		MQ:    mq.New(cfg.MQ),
		Graph: graphql.New(cfg.GraphQL),
		HTTP:  http.New(),
		RPC:   rpc.New(cfg.Client, isRpcServer),
	}
}

func (m *MetaResource) Close() {
	m.Async.Close()
	m.DB.Close()
	m.Cache.Close()
	m.MQ.Close()
	m.DB.Close()
	m.Graph.Close()
	m.HTTP.Close()
	m.RPC.Close()
}
