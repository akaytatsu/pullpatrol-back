package git

type GitInterface interface {
	GetUserData(userCode string) (dataStruct any, err error)
	ProccessWebhook(payload []byte) (dataStruct any, err error)
	Driver() string
}
