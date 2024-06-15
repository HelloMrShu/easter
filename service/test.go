package service

import "fmt"

type TestService struct {
}

func NewTestService() *TestService {
	return &TestService{}
}

func (t *TestService) GetInfo(name string) string {
	return fmt.Sprintf("hello %s!", name)
}
