package simple_store

import (
	"fmt"
	"github.com/iiran/lltt/pkg/helper"
	"github.com/iiran/lltt/pkg/logger"
	"github.com/iiran/lltt/pkg/model"
	"github.com/iiran/lltt/pkg/setting"
	"time"
)

const _KILL_ON_FULL = 1000

// one user one session
type SessionStore struct {
	seed string
	max  int64
	cur  int64
	db   map[string]model.UserSession // session-id -> user-session
}

func (ss *SessionStore) Exist(key string) bool {
	return ss.Get(key) != nil
}

func (ss *SessionStore) Get(key string) *model.UserSessionData {
	userss, exist := ss.db[key]
	if !exist {
		return nil
	}
	// lazy clean expired session
	if userss.ExpireTime.Before(time.Now()) {
		delete(ss.db, key)
		return nil
	}
	return &userss.Data
}

func (ss *SessionStore) SetHours(data model.UserSessionData, lifeHour int64) (key string) {
	now := time.Now()

	userss := model.UserSession{
		Data:       data,
		CreateTime: now,
		ExpireTime: now.Add(time.Duration(helper.HourToNano(lifeHour))),
	}

	if ss.cur > ss.max {
		ss.DeleteRandomN(_KILL_ON_FULL)
	}
	key = generateSessionID()
	ss.db[key] = userss
	ss.cur++
	logger.Info(fmt.Sprintf(`create session success. session-id = %s`, key))
	return
}

func (ss *SessionStore) Delete(key string) {
	delete(ss.db, key)
}

func (ss *SessionStore) ExtendHour(key string, hours int64) {
	if hours <= 0 {
		return
	}
	userSession, exist := ss.db[key]
	if !exist {
		return
	}
	userSession.ExpireTime.Add(time.Duration(helper.HourToNano(hours)))
	ss.db[key] = userSession
}

func (ss *SessionStore) Reset(key string, hours int64) bool {
	if hours <= 0 {
		return false
	}
	userSession, exist := ss.db[key]
	if !exist {
		return false
	}
	now := time.Now()
	userSession.CreateTime = now
	userSession.ExpireTime = now.Add(time.Duration(helper.HourToNano(hours)))
	ss.db[key] = userSession
	return true
}

func (ss *SessionStore) DeleteRandomN(n int64) {
	if n <= 0 {
		return
	}
	var cur int64
	for k := range ss.db {
		delete(ss.db, k)
		cur++
		if cur >= n {
			break
		}
	}
}

func NewSessionStore(seed string, size int64) *SessionStore {
	return &SessionStore{db: map[string]model.UserSession{}, seed: seed, max: size}
}

func Setup(cfg *setting.ServerConfigSession) {
	store = NewSessionStore(cfg.Secret, cfg.Count)
}

var store *SessionStore

func Get(key string) *model.UserSessionData {
	return store.Get(key)
}

func Delete(key string) {
	store.Delete(key)
}

func SetHours(data model.UserSessionData, hours int64) (sessionID string) {
	sessionID = store.SetHours(data, hours)
	return
}
