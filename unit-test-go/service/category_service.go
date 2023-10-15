package service

import(
	"unit-test-go/repository"
	"unit-test-go/entity"
	"errors"
)

type CategoryService struct{
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error){
	category := service.Repository.FindById(id)
	if category==nil{
		return nil, errors.New("category not found")
	}else{
		return category, nil
	}

}