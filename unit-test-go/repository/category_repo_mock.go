package repository

import(
	"unit-test-go/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepoMock struct{
	Mock mock.Mock
}

func (repository *CategoryRepoMock) FindById(id string) *entity.Category{
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil{
		return nil
	}else{
		category :=  arguments.Get(0).(entity.Category)
		return &category
	}
}