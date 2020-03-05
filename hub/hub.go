package hub

import "github.com/zu1k/she/hub/route"

// Start 启动交互接口
func Start() {
	route.Start("127.0.0.1:19090", "")
}
