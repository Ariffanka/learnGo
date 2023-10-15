package repository

import(
	"unit-test-go/entity"
)

type CategoryRepository interface{
		FindById(id string) *entity.Category
}