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
 @Time    : 2025/7/2 -- 14:23
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: seele /assembler.go
*/

package matrix

import "context"

type Assembler[ent any, info any] interface {
	Assemble(context.Context, ent) *info
	Disassemble(context.Context, info) *ent
}

type AssemblerUniversal[Info any] interface {
	Assemble(ctx context.Context, ent ...any) *Info
	Disassemble(context.Context, Info) []any
}
