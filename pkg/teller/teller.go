package teller

type TellerAPI struct {
	appId string
	token string
}

func Init(appId string, token string) *TellerAPI {
	return &TellerAPI{
		appId: appId,
		token: token,
	}
}
