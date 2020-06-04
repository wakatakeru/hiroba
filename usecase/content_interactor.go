package usecase

import "github.com/wakatakeru/hiroba-content-api/domain"

type ContentInteractor struct {
	ContentRepository ContentRepository
}

func NewContentInteractor(contentRepository ContentRepository) ContentInteractor {
	contentInteractor := ContentInteractor{ContentRepository: contentRepository}
	return contentInteractor
}

func (interactor *ContentInteractor) Add(c domain.Content) (id int, err error) {
	id, err = interactor.ContentRepository.Store(c)
	return
}

func (interactor *ContentInteractor) Update(c domain.Content) (count int, err error) {
	count, err = interactor.ContentRepository.Update(c)
	return
}

func (interactor *ContentInteractor) Content(id int) (content domain.Content, err error) {
	content, err = interactor.ContentRepository.FindByID(id)
	return
}

func (interactor *ContentInteractor) SiteContents(siteID int) (contents domain.Contents, err error) {
	contents, err = interactor.ContentRepository.FindBySiteID(siteID)
	return
}

func (interactor *ContentInteractor) UserContents(userID int) (contents domain.Contents, err error) {
	contents, err = interactor.ContentRepository.FindByUserID(userID)
	return
}
