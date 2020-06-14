package database

import (
	"github.com/wakatakeru/hiroba/apis/content/domain"
)

type ContentRepository struct {
	SqlHandler
}

func NewContenRepository(sqlHandler SqlHandler) ContentRepository {
	contentRepository := ContentRepository{SqlHandler: sqlHandler}
	return contentRepository
}

func (repo *ContentRepository) Store(c domain.Content) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO contents (site_id, user_id, title, body) VALUES (?,?,?,?)",
		c.SiteID,
		c.UserID,
		c.Title,
		c.Body,
	)

	if err != nil {
		return
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *ContentRepository) Update(c domain.Content) (count int, err error) {
	result, err := repo.Execute(
		"UPDATE contents SET site_id=?, user_id=?, title=?, body=?",
		c.SiteID,
		c.UserID,
		c.Title,
		c.Body,
	)

	if err != nil {
		return
	}

	count64, err := result.RowsAffected()
	if err != nil {
		return
	}

	count = int(count64)
	return
}

func (repo *ContentRepository) FindByID(id int) (content domain.Content, err error) {
	row, err := repo.Query(
		"SELECT id, site_id, user_id, title, body contents WHERE id=?",
		id,
	)
	defer row.Close()
	if err != nil {
		return
	}

	var siteID int
	var userID int
	var title string
	var body string

	row.Next()
	err = row.Scan(&id, &siteID, &userID, &title, &body)
	if err != nil {
		return
	}

	content.ID = id
	content.SiteID = siteID
	content.UserID = userID
	content.Title = title
	content.Body = body
	return
}

func (repo *ContentRepository) FindBySiteID(siteID int) (contents domain.Contents, err error) {
	row, err := repo.Query(
		"SELECT id, site_id, user_id, title, body contents WHERE site_id=?",
		siteID,
	)
	defer row.Close()
	if err != nil {
		return
	}

	for row.Next() {

		var id int
		var siteID int
		var userID int
		var title string
		var body string

		if err = row.Scan(&id, &siteID, &userID, &title, &body); err != nil {
			continue
		}

		content := domain.Content{
			ID:     id,
			SiteID: siteID,
			UserID: userID,
			Title:  title,
			Body:   body,
		}

		contents = append(contents, content)
	}

	return
}

func (repo *ContentRepository) FindByUserID(userID int) (contents domain.Contents, err error) {
	row, err := repo.Query(
		"SELECT id, site_id, user_id, title, body contents WHERE user_id=?",
		userID,
	)
	defer row.Close()
	if err != nil {
		return
	}

	for row.Next() {

		var id int
		var siteID int
		var userID int
		var title string
		var body string

		if err = row.Scan(&id, &siteID, &userID, &title, &body); err != nil {
			continue
		}

		content := domain.Content{
			ID:     id,
			SiteID: siteID,
			UserID: userID,
			Title:  title,
			Body:   body,
		}

		contents = append(contents, content)
	}

	return
}
