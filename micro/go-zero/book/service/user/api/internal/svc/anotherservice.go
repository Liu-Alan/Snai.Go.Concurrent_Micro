package svc

import "github.com/zeromicro/go-zero/core/stringx"

type AnotherService struct{}

func NewAnotherService() *AnotherService {
	return &AnotherService{}
}

func (s *AnotherService) GetToken() string {
	return stringx.Rand()
}
