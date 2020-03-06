package netease

import "github.com/zu1k/she/source"

type netease struct {
}

func init() {
	source.Register("netease", newNetease)
}

// GetName return 126/163 name
func (n *netease) GetName() string {
	return "Netease"
}

// Search return result slice from source 126/163
func (n *netease) Search(key interface{}) (result []source.Result) {
	return nil
}

func newNetease(info interface{}) source.Source {
	return &netease{}
}
