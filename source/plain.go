package source

type plaintext struct {
}

func init() {
	register("plaintext", newPlain)
}

// GetName return plaintext name
func (p *plaintext) GetName() string {
	return "Plain"
}

// Search return result slice from source plaintext
func (p *plaintext) Search(key interface{}) (result []Result) {
	return nil
}

func newPlain(info interface{}) Source {
	return &plaintext{}
}
