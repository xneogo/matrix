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
 @Time    : 2024/10/17 -- 18:07
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2024 亓官竹
 @Description: server.go
*/

package mserv

import (
	"context"
	"github.com/xneogo/matrix/mconfig"
)

type ServerSessionProxy[servInfo any] interface {
	RegisterProxy[servInfo]

	// ------ infos defalult get ------

	Name(ctx context.Context) string                    // eg: 服务 user | account | markettask | introduction | marketapi | salesopapi
	Group(ctx context.Context) string                   // eg: 服务组 base | market | api | opapi | sales
	ServiceKey(ctx context.Context) string              // eg: 服务名 base/user | market/markettask
	Lane(ctx context.Context) string                    // 泳道
	Region(ctx context.Context) string                  // region
	Ip(ctx context.Context) string                      // ip
	ID(ctx context.Context) int                         // e.g: 1
	ServInfos(ctx context.Context) map[string]*servInfo // serv infos
	RegInfos() map[string]string                        // reg infos
	FullName(ctx context.Context) string                // eg: 服务全名 trade/points1 | base/user1
	ConfigCenter(ctx context.Context) mconfig.ConfigCenter

	// ------ actions to operate server ------

	Startup(ctx context.Context, args interface{}) error      // 启动服务
	ApplyPreFns(ctx context.Context, args interface{}) error  // 执行项目服务实例初始化之前的准备工作
	ApplyPostFns(ctx context.Context, args interface{}) error // 执行项目服务实例初始化之后的准备 or 清理工作
	Shutdown(ctx context.Context) error                       // 关停服务
	AppendShutdownCallback(context.Context, func())           // 添加关停回调函数

	// ------ statement of server ------

	// Offline 服务状态 true offline; false online
	Offline(ctx context.Context) bool
	// RunningLocal return true if server is local running
	RunningLocal(ctx context.Context) bool

	// WithControlLaneInfo wrap context with service context info, such as lane
	WithControlLaneInfo(ctx context.Context) context.Context
}

type ZProcessor interface {
	Init() error
	Driver() (string, interface{})
}

type UUID interface {
	Gen() (string, error)
	GenSha1() (string, error)
	GenMd5() (string, error)
}

type RegisterProxy[SvcInfo any] interface {
	// Register key is processor to ServInfo; dir is service name like account | user |
	Register(ctx context.Context, svcMap map[string]*SvcInfo, dir string) error
	// RegisterService SvcLocRegType Service by default
	RegisterService(ctx context.Context, svcMap map[string]*SvcInfo, dir string, cross bool) error
	// RegisterCrossDCService SvcLocRegType Service by default
	RegisterCrossDCService(ctx context.Context, svcMap map[string]*SvcInfo, dir string) error
}

// EtcdLocker
// 默认的锁，局部分布式锁，各个服务之间独立不共享
type EtcdLocker interface {
	// Lock 获取到lock立即返回，否则block直到获取到
	Lock(ctx context.Context, name string) error
	// Unlock 没有lock的情况下unlock，程序会直接panic
	Unlock(ctx context.Context, name string) error
	// TryLock 立即返回，如果获取到lock返回true，否则返回false
	TryLock(ctx context.Context, name string) (bool, error)
}

// EtcdGlobalLocker
// 全局分布式锁，全局只有一个，需要特殊加global说明
type EtcdGlobalLocker interface {
	LockGlobal(ctx context.Context, name string) error
	UnlockGlobal(ctx context.Context, name string) error
	TryLockGlobal(ctx context.Context, name string) (bool, error)
}
