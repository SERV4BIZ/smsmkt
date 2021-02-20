package smsmkt

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/SERV4BIZ/gfp/jsons"
)

// Send is send sms
func (me *SMSMKT) Send(txtSenderName string, txtPhone string, txtMessage string) (*jsons.JSONObject, error) {
	jsoData := jsons.ObjectNew()
	jsoData.PutString("message", txtMessage)
	jsoData.PutString("phone", txtPhone)
	jsoData.PutString("sender", txtSenderName)

	url := "https://portal-otp.smsmkt.com/api/send-message"
	method := "POST"
	payload := strings.NewReader(jsoData.ToString())

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
