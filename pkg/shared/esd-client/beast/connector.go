package beast

type Connector interface {
	RunJob() (string, error)
	GetLifecycleStage() (string, error)
	GetRuntimeInfo() (string, error)
	GetConfiguration() (string, error)
	GetLogs() (string, error)
}

type connector struct {
	url      string
	codeRoot string
}

func NewConnector() Connector {
	return &connector{}
}
