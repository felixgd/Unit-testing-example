package service

type ServiceInterface interface {
	ReturnText(s string) (string, error)
}
