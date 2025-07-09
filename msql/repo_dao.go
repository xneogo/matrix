/*
 *  ┏┓      ┏┓
 *┏━┛┻━━━━━━┛┻┓
 *┃　　　━　　  ┃
 *┃   ┳┛ ┗┳   ┃
 *┃           ┃
 *┃     ┻     ┃
 *┗━━━┓     ┏━┛
 *　　 ┃　　　┃神兽保佑
 *　　 ┃　　　┃代码无BUG！
 *　　 ┃　　　┗━━━┓
 *　　 ┃         ┣┓
 *　　 ┃         ┏┛
 *　　 ┗━┓┓┏━━┳┓┏┛
 *　　   ┃┫┫  ┃┫┫
 *      ┗┻┛　 ┗┻┛
 @Time    : 2024/11/11 -- 12:03
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: repo_dao.go
*/

package msql

import "context"

type RepoConverter[Data any, Entity any] interface {
	ToEntity(ctx context.Context, data Data) *Entity
	MultiToEntity(ctx context.Context, ds []Data) []*Entity
}

// DaoModel
// interactive between DataObj and db
// reflect db instance rows to DataObj
type DaoModel[DObj any] interface {
	Init(cons SqlConstructor, tableName func() string, omits func() []string, b BindFunc)
	TableName() string
	Omits() []string
	GetScanner() Scanner
	GetBuilder() Builder

	SelectOne(ctx context.Context, db XDB, where map[string]interface{}) (DObj, error)
	SelectMulti(ctx context.Context, db XDB, where map[string]interface{}) ([]DObj, error)
	Insert(ctx context.Context, db XDB, data []map[string]interface{}) (int64, error)
	Upsert(ctx context.Context, db XDB, data map[string]interface{}) (int64, error)
	Update(ctx context.Context, db XDB, where, data map[string]interface{}) (int64, error)
	Delete(ctx context.Context, db XDB, where map[string]interface{}) (int64, error)
	CountOf(ctx context.Context, db XDB, where map[string]interface{}) (count int, err error)
}
