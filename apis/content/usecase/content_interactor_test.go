package usecase

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	domain "github.com/wakatakeru/hiroba/apis/content/domain"
)

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)

	var expectedContent domain.Content
	var err error
	var expectedID int

	mockRepository := NewMockContentRepository(ctrl)
	mockRepository.EXPECT().Store(expectedContent).Return(expectedID, err)

	contentInteractor := NewContentInteractor(mockRepository)
	id, err := contentInteractor.Add(expectedContent)

	if err != nil {
		t.Error("Add() is not same as expected")
	}

	if !reflect.DeepEqual(expectedID, id) {
		t.Error("Add() is not same as expected")
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)

	var expectedContent domain.Content
	var err error
	var expectedCount int

	mockRepository := NewMockContentRepository(ctrl)
	mockRepository.EXPECT().Store(expectedContent).Return(expectedCount, err)

	contentInteractor := NewContentInteractor(mockRepository)
	count, err := contentInteractor.Add(expectedContent)

	if err != nil {
		t.Error("Update() is not same as expected")
	}

	if !reflect.DeepEqual(expectedCount, count) {
		t.Error("Update() is not same as expected")
	}
}

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

func TestSiteContents(t *testing.T) {
	ctrl := gomock.NewController(t)

	var expectedContents domain.Contents
	var err error
	var siteID int

	mockRepository := NewMockContentRepository(ctrl)
	mockRepository.EXPECT().FindBySiteID(siteID).Return(expectedContents, err)

	contentInteractor := NewContentInteractor(mockRepository)
	resultContents, err := contentInteractor.SiteContents(siteID)

	if err != nil {
		t.Error("SiteContents() is not same as expected")
	}

	if !reflect.DeepEqual(resultContents, expectedContents) {
		t.Error("SiteContents() is not same as expected")
	}
}

func TestUserContents(t *testing.T) {
	ctrl := gomock.NewController(t)

	var expectedContents domain.Contents
	var err error
	var userID int

	mockRepository := NewMockContentRepository(ctrl)
	mockRepository.EXPECT().FindByUserID(userID).Return(expectedContents, err)

	contentInteractor := NewContentInteractor(mockRepository)
	resultContents, err := contentInteractor.UserContents(userID)

	if err != nil {
		t.Error("UserContents() is not same as expected")
	}

	if !reflect.DeepEqual(resultContents, expectedContents) {
		t.Error("UserContents() is not same as expected")
	}
}
