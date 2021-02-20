package smsmkt

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/SERV4BIZ/gfp/jsons"
)

// GetCredit is get balance of credit
func (me *SMSMKT) GetCredit() (*jsons.JSONObject, error) {
	url := "https://portal-otp.smsmkt.com/api/get-credit"
	method := "GET"
	payload := strings.NewReader("{}")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api_key", me.APIKey)
	req.Header.Add("secret_key", me.SecretKey)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	jsoRes, errRes := jsons.ObjectString(string(body))
	if errRes != nil {
		return nil, err
	}

	return jsoRes, nil
}
