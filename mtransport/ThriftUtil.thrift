namespace go util.thriftutil
namespace php util.thriftutil
namespace py util.thriftutil
namespace cpp util.thriftutil
namespace java util.thriftutil

struct ErrInfo {
  1: required i32 code
  2: required string msg
}

struct Head {
    1: i64 uid
    2: i32 source
	3: string ip
	4: string region
	5: i32 dt
	6: string unionid
    7: string did
    8: i32 zone
    9: string zone_name
}

struct Control {
    1: optional Route route
    2: i64 ct
    3: i64 et
    4: optional Endpoint caller
}

struct Route {
    1: string group
}

struct Endpoint {
    // server name
    1: string sname
    // server id
    2: string sid
    // rpc name
    3: string method
}

struct Context {
	1: optional Head head
	2: optional map<string, string> spanctx
	3: optional Control control
}
