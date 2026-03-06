package domain

import "context"

type DepartementDomain struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type DepartementsRepository interface {
	FindAllDepartements(ctx context.Context) (result []DepartementDomain, err error)
	FindDepartementById(ctx context.Context, departementId string) (result DepartementDomain, err error)
	SaveDepartement(ctx context.Context, departement DepartementDomain) error
	DeleteDepartement(ctx context.Context, departementId string) error
}

type DepartementService interface {
	FindAllDepartements(ctx context.Context) (result []DepartementDomain, err error)
	FindDepartementById(ctx context.Context, departementId string) (result DepartementDomain, err error)
	CreateDepartement(ctx context.Context, departement DepartementDomain) error
	DeleteDepartement(ctx context.Context, departementId string) error
}
