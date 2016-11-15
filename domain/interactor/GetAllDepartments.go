package interactor

import (
	data "Rein/data/repository"
	"Rein/domain"
	"Rein/domain/repository"
)

// GetAllDepartments Use case for create a new Department
type GetAllDepartments struct {
	UseCase
	departmentRepository repository.DepartmentRepository
}

// NewGetAllDepartments Constructor for CreateDepartment UseCase
func NewGetAllDepartments() *GetAllDepartments {
	return &GetAllDepartments{departmentRepository: data.NewDepartmentDataRepository()}
}

// Execute The use case
func (u *GetAllDepartments) Execute(ch chan []domain.Department) {
	/*c := u.departmentRepository.AddDepartment(u.department)
	department := <-c
	fmt.Println(department.ToString())

	ch <- department*/

	ch <- <-u.departmentRepository.GetDepartments()
}