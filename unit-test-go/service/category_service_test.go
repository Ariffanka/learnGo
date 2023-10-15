package service

import(
	"testing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"unit-test-go/repository"
	"unit-test-go/entity"
)

var catRepo= &repository.CategoryRepoMock{Mock: mock.Mock{}}
var catService= CategoryService{Repository: catRepo}

func TestCategoryServiceTest(t *testing.T){

	catRepo.Mock.On("FindById", "1").Return(nil)
	category, err := catService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryServiceSucces(t *testing.T){
	category:= entity.Category{
		Id: "1",
		Name: "Laptop",
	}

	catRepo.Mock.On("FindById", "2").Return(category)

	result, err := catService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}