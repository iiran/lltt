package service

import (
	"database/sql"
	"github.com/iiran/lltt/pkg/core/errors"
	"github.com/iiran/lltt/pkg/db"
	"github.com/iiran/lltt/pkg/model"
)

func CreateReply(mould model.ReplyMould) (err error) {
	var (
		q = `insert into reply (content, author, reply_type, reply_to) values (?, ?, ?, ?)`
	)
	if _, err = db.Exec(q, mould.Content, mould.ReplyFrom, mould.ReplyType, mould.ReplyTo); err != nil {
		return errors.ErrAppend(err, errors.DB_INSERT_ERR)
	}
	return nil
}

func GetReply(replyID int64) (repl model.Reply, err error) {
	var (
		rows *sql.Rows
		q    = `select content, created_at, author, reply_to from reply where id = ?`
	)
	if rows, err = db.Query(q, replyID); err != nil {
		return repl, errors.ErrAppend(err, errors.DB_SELECT_ERR)
	}
	defer rows.Close()
	if !rows.Next() {
		return repl, errors.ErrAppend(err, errors.NOT_EXIST)
	}
	if err = rows.Scan(&repl.Content, &repl.CreatedAt, &repl.Author, &repl.ReplyTo); err != nil {
		return repl, errors.ErrAppend(err, errors.DB_SCAN_ERR)
	}
	repl.ID = replyID
	return
}

func GetRepliesFrom(replyType int, from int64, offset int64, limit int64) (repls []model.Reply, err error) {
	var (
		rows *sql.Rows
		q    = `select id, content, created_at, reply_to from reply where reply_type = ? and author = ? limit ? offset ?`
	)
	if rows, err = db.Query(q, replyType, from, limit, offset); err != nil {
		return nil, errors.ErrAppend(err, errors.DB_SELECT_ERR)
	}
	defer rows.Close()
	for rows.Next() {
		repl := model.Reply{}
		if err = rows.Scan(&repl.ID, &repl.Content, &repl.CreatedAt, &repl.ReplyTo); err != nil {
			return repls, errors.ErrAppend(err, errors.DB_SCAN_ERR)
		}
		repl.Author = from
		repls = append(repls, repl)
	}
	return
}

func GetRepliesTo(replyType int, to int64, offset int64, limit int64) (repls []model.Reply, err error) {
	var (
		rows *sql.Rows
		q    = `select id, content, created_at, author from reply where reply_type = ? and reply_to = ? limit ? offset ?`
	)
	if rows, err = db.Query(q, replyType, to, limit, offset); err != nil {
		return nil, errors.ErrAppend(err, errors.DB_SELECT_ERR)
	}
	defer rows.Close()
	for rows.Next() {
		repl := model.Reply{}
		if err = rows.Scan(&repl.ID, &repl.Content, &repl.CreatedAt, &repl.Author); err != nil {
			return repls, errors.ErrAppend(err, errors.DB_SCAN_ERR)
		}
		repl.ReplyTo = to
		repls = append(repls, repl)
	}
	return
}
