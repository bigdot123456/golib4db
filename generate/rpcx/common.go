package rpcx

import "github.com/bigdot123456/golib4db/generate"

func init() {
	generate.RegisterLayouter("rpcxclient", &layclient{})
	generate.RegisterLayouter("rpcx", &lay{})
}
