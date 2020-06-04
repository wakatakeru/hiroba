package usecase

import (
	"github.com/wakatakeru/hiroba-content-api/domain"
)

type ContentRepository interface {
	Store(domain.Content) (int, error)
	Update(domain.Content) (int, error)
	FindByID(int) (domain.Content, error)
	FindBySiteID(int) (domain.Contents, error)
	FindByUserID(int) (domain.Contents, error)
}
