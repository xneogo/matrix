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
 @Time    : 2025/7/2 -- 18:29
 @Author  : 亓官竹 ❤️ MONEY
 @Copyright 2025 亓官竹
 @Description: mentity mentity/governor.go
*/

package mentity

import (
	"fmt"
	"strings"
)

type ProcessorType string

type ServInfo struct {
	Type       ProcessorType `json:"type"`
	Addr       string        `json:"addr"`
	ServId     int           `json:"-"`
	ServiceKey string        `json:"serviceKey"` // eg: base/user
	// Processor string    `json:"processor"`
}

func (m *ServInfo) String() string {
	return fmt.Sprintf("%s://%s", m.Type, m.Addr)
}

type RegData struct {
	ServMap map[string]*ServInfo `json:"serv_map"`
	Lane    *string              `json:"lane"`
}

func NewRegData(servM map[string]*ServInfo, lane string) *RegData {
	return &RegData{
		ServMap: servM,
		Lane:    &lane,
	}
}

func (r *RegData) GetLane() (string, bool) {
	if r.Lane == nil {
		return "", false
	}
	return *r.Lane, true
}

func (r *RegData) String() string {
	var procs []string
	for k, v := range r.ServMap {
		procs = append(procs, fmt.Sprintf("%s@%s", v, k))
	}
	return strings.Join(procs, "|")
}
