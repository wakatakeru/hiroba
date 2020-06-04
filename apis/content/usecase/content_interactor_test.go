package usecase

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	domain "github.com/wakatakeru/hiroba/apis/content/domain"
)

func TestContent(t *testing.T) {
	ctrl := gomock.NewController(t)

	var expectedContent domain.Content
	var err error
	var id int

	mockRepository := NewMockContentRepository(ctrl)
	mockRepository.EXPECT().FindByID(id).Return(expectedContent, err)

	contentInteractor := NewContentInteractor(mockRepository)
	resultContent, err := contentInteractor.Content(id)

	if err != nil {
		t.Error("Content() is not same as expected")
	}

	if !reflect.DeepEqual(resultContent, expectedContent) {
		t.Error("Content() is not same as expected")
	}
}
