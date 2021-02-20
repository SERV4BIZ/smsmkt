package smsmkt

// Factory is create smsmkt object from api and secret key
func Factory(txtAPIKey string, txtSecretKey string) *SMSMKT {
	item := new(SMSMKT)
	item.APIKey = txtAPIKey
	item.SecretKey = txtSecretKey
	return item
}
