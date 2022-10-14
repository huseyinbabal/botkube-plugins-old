package botkubeplugin

type Source interface {
	Consume(ch chan interface{}) error
}
