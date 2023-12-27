package boxer

type Connector interface {
	Token
	Claim
}

type User interface {
	AddUser() (string, error)
	RemoveUser() (string, error)
}

type Token interface {
	GetToken() (string, error)
}

type Claim interface {
	GetClaim(user string, provider string) (string, error)
	AddClaim(user string, provider string, claims []string) (string, error)
	RemoveClaim(user string, provider string, claims []string) (string, error)
}

// connector is an implementation of the Connector interface.
type connector struct {
	tokenUrl string
	claimUrl string
	auth     ExternalToken
	retries  int
}

type Input struct {
	TokenUrl string
	ClaimUrl string
	Auth     ExternalToken
	Retries  int
}

// NewConnector creates a new Connector instance with the provided configuration.
func NewConnector(i Input) Connector {
	return &connector{
		tokenUrl: i.TokenUrl,
		claimUrl: i.ClaimUrl,
		auth:     i.Auth,
		retries:  i.Retries,
	}
}
