package core

type Credential struct {
	SecretKey string
}

func NewCredentials(secretKey string) *Credential {
	return &Credential{secretKey}
}
