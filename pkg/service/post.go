package service

import (
	"database/sql"
	"github.com/iiran/lltt/pkg/core/errors"
	"github.com/iiran/lltt/pkg/db"
	"github.com/iiran/lltt/pkg/model"
)

func ListPosts(offset int64, limit int64) (posts []model.Post, err error) {
	if limit <= 0 {
		return make([]model.Post, 0), nil
	}
	var (
		rows *sql.Rows
		q    = `select id, title, describe, content, created_at, creator, author, tag_name from posts limit ? offset ?`
	)
	if rows, err = db.Query(`pqs`, q, limit, offset); err != nil {
		return nil, errors.ErrAppend(err, errors.DB_SELECT_ERR)
	}
	defer rows.Close()
	posts = make([]model.Post, 0)
	for rows.Next() {
		post := model.Post{}
		var _tagName sql.NullString
		if err = rows.Scan(&post.ID, &post.Title, &post.Describe, &post.Content, &post.CreatedAt, &post.Author, &post.Creator, &_tagName); err != nil {
			return posts, errors.ErrAppend(err, errors.DB_SCAN_ERR)
		}
		post.Tag = _tagName.String
		posts = append(posts, post)
	}
	return
}

func ListUserPosts(author int64, offset int64, limit int64) (posts []model.Post, err error) {
	if limit <= 0 {
		return make([]model.Post, 0), nil
	}
	var (
		rows *sql.Rows
		q    = `select id, title, describe, content, created_at, creator, tag_name from posts where author = ? limit ? offset ?`
	)
	if rows, err = db.Query(`pqs`, q, author, limit, offset); err != nil {
		return nil, errors.ErrAppend(err, errors.DB_SELECT_ERR)
	}
	defer rows.Close()
	posts = make([]model.Post, 0)
	for rows.Next() {
		post := model.Post{}
		var _tagName sql.NullString
		if err = rows.Scan(&post.ID, &post.Title, &post.Describe, &post.Content, &post.CreatedAt, &post.Author, &post.Creator, &_tagName); err != nil {
			return posts, errors.ErrAppend(err, errors.DB_SCAN_ERR)
		}
		post.Tag = _tagName.String
		post.Author = author
		posts = append(posts, post)
	}
	return
}

func QueryPost(id int64) (post model.Post, err error) {
	var (
		rows     *sql.Rows
		_tagName sql.NullString
		q        = `select title, describe, content, created_at, creator, author, tag_name from posts where id = ?`
	)
	if rows, err = db.Query(`pqs`, q, id); err != nil {
		return post, errors.ErrAppend(err, errors.DB_SELECT_ERR)
	}
	defer rows.Close()
	if !rows.Next() {
		return post, errors.GetErr(errors.NOT_EXIST)
	}
	if err = rows.Scan(&post.Title, &post.Describe, &post.Content, &post.CreatedAt, &post.Creator, &post.Author, &_tagName); err != nil {
		return post, errors.ErrAppend(err, errors.DB_SCAN_ERR)
	}
	post.Tag = _tagName.String
	post.ID = id
	return
}

func CreatePost(creator int64, data model.PostMould) (created model.PostMould, err error) {
	var (
		q = `insert into posts (title, describe, content, creator, author) values (?, ?, ?, ?, ?)`
	)
	if _, err = db.Exec(`pqs`, q, data.Title, data.Describe, data.Content, creator, creator); err != nil {
		return created, errors.ErrAppend(err, errors.DB_INSERT_ERR)
	}
	created.Content, created.Describe, created.Title = data.Content, data.Describe, data.Title
	return created, nil
}
