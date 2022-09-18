package main

import (
	"errors"
	"fmt"

	//"github.com/louistwiice/go/basicwithent/entity"
	//"github.com/louistwiice/go/basicwithent/mocks"
	"github.com/stretchr/testify/mock"
)

type mockObject struct{
	mock.Mock
}

type Service interface {
    RequestForInformation(request string) (string, string, error)
}
  
type Business struct {
    service Service
}
 
func (b Business) doGoodThings(theThing string) (string, error) {
 
    result, _, _ := b.service.RequestForInformation(theThing)
    if len(result) == 0 {
        return "", fmt.Errorf("Bad result")
    }
    return result, nil
}

func (m mockObject) RequestForInformation(request string) (string, string, error) {
    args := m.Called(request)
	fmt.Println("1111111 = ", args)
	fmt.Println("2222 = ", args.String(0))
	fmt.Println("3333 = ", args.Get(1).(string))
    return args.String(0), args.String(0), args.Error(1)
}

func main() {
	serviceMock := mockObject{}
	serviceMock.On("RequestForInformation", "GoodRequest").Return("YEE","OK", errors.New("ss"))

    b := Business{service: serviceMock}
    result, err := b.doGoodThings("GoodRequest")
	fmt.Println("====== ", result)
	fmt.Println("====== ", err)
}
