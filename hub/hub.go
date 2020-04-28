package hub

import "github.com/zu1k/she/hub/route"

// Start 启动交互接口
func Start(addr, secret string) {
	route.Start(addr, secret)
}
