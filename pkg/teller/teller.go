package teller

import (
	"net/http"
)

type TellerAPI struct {
	appId    string
	token    string
	baseUrl  string
	certPath string
	keyPath  string
}

func Init(appId string, token string, baseUrl string, certPath string, keyPath string) *TellerAPI {
	return &TellerAPI{
		appId:    appId,
		token:    token,
		baseUrl:  baseUrl,
		certPath: certPath,
		keyPath:  keyPath,
	}
}

func (t *TellerAPI) GetAccounts() error {
	return nil
}

func (t *TellerAPI) DoTellerHttp(method string, path string) error {
    url := t.baseUrl + path

    _, err := http.NewRequest(method, url, nil)
    if err != nil {
        return err
    }

    return nil
}
