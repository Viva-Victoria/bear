package bear

type Logger interface {
	Warn(message string, err error)
	Error(err error)
}
