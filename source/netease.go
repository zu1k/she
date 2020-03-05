package source

type netease struct {
}

func init() {
	register("netease", newNetease)
}

// GetName return 126/163 name
func (n *netease) GetName() string {
	return "QQGroup"
}

// Search return result slice from source 126/163
func (n *netease) Search(key interface{}) (result []Result) {
	return nil
}

func newNetease(info interface{}) Source {
	return &netease{}
}
