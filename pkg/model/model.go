package model

import "time"

type User struct {
	ID          int64
	Username    string
	Displayname string
	Email       string
	CreatedAt   time.Time
	Score       int64
	Weight      int64
}

type UserMould struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserLoginMould struct {
	Username string `json:"username"`
}

type ReplyMould struct {
	ReplyType int32
	ReplyFrom int64
	ReplyTo   int64
	Content   string
}

type Reply struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Author    int64  `json:"author"`
	ReplyTo   int64  `json:"reply_to"`
}

const (
	ReplyTypeUser  = 1
	ReplyTypePost  = 2
	ReplyTypeReply = 3
)

type UserReplyMould struct {
	Content string `json:"content"`
}

type PostReplyMould struct {
	Content string `json:"content"`
}

type ReplyReplyMould struct {
	Content string `json:"content"`
}

func CreateUserReplyMould(from int64, to int64, data UserReplyMould) ReplyMould {
	return ReplyMould{
		ReplyType: 1,
		ReplyFrom: from,
		ReplyTo:   to,
		Content:   data.Content,
	}
}

func CreatePostReplyMould(from int64, to int64, data PostReplyMould) ReplyMould {
	return ReplyMould{
		ReplyType: 2,
		ReplyFrom: from,
		ReplyTo:   to,
		Content:   data.Content,
	}
}

func CreateReplyReplyMould(from int64, to int64, data ReplyReplyMould) ReplyMould {
	return ReplyMould{
		ReplyType: 3,
		ReplyFrom: from,
		ReplyTo:   to,
		Content:   data.Content,
	}
}

type Post struct {
	ID        int64
	Title     string
	Describe  string
	Content   string
	CreatedAt string
	Author    int64
	Creator   int64
	Tag       string
}

type PostMould struct {
	Title    string `json:"title"`
	Describe string `json:"describe"`
	Content  string `json:"content"`
}

type UserSessionData struct {
	UserID int64
}

type UserSession struct {
	Data       UserSessionData
	CreateTime time.Time
	ExpireTime time.Time
}
