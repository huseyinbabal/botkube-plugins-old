package botkubeplugin

type Executor interface {
	Execute(command string) (string, error)
}
