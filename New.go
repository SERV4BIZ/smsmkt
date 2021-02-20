package smsmkt

// New is create smsmkt object from api and secret key
func New(txtAPIKey string, txtSecretKey string) *SMSMKT {
	return Factory(txtAPIKey, txtSecretKey)
}
