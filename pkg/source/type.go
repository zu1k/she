package source

type Type int

const (
	BleveIndex = iota
	QQGroup
	PlainText
)

func (t Type) String() string {
	switch t {
	case QQGroup:
		return "QQGroup"
	case PlainText:
		return "PlainText"
	case BleveIndex:
		return "BleveIndex"
	default:
		return "NotKnown"
	}
}
