package session

import (
	"fmt"
	"github.com/wjpxxx/letgo/cache"
	"github.com/wjpxxx/letgo/cache/filecache"
	"github.com/wjpxxx/letgo/cache/icache"
	"github.com/wjpxxx/letgo/encry"
	"github.com/wjpxxx/letgo/file"
	"github.com/wjpxxx/letgo/lib"
	"sync"
	"time"
)

//Session
type Session struct {
	cache  icache.ICacher
	cookie Cookier
}

//Set
func (s *Session) Set(key string, value interface{}) bool {
	return s.cache.Set(s.Name(key), value, int64(config.Expire))
}

//Get
func (s *Session) Get(key string, value interface{}) bool {
	return s.cache.Get(s.Name(key), value)
}

//Del
func (s *Session) Del(key string) bool {
	return s.cache.Del(s.Name(key))
}

//FlushDB
func (s *Session) FlushDB() bool {
	return s.cache.FlushDB()
}

//sessionMutex
var sessionMutex sync.Mutex

//CreateSessionID
func (s *Session) CreateSessionID() string {
	sessionMutex.Lock()
	defer sessionMutex.Unlock()
	return encry.Base64Encode(encry.Hmac(fmt.Sprintf("%d", time.Now().UnixNano()), lib.RandChar(20)))
}
//SessionKey
func (s *Session) SessionKey() string {
	return config.Name
}
//SessionID
func (s *Session) SessionID() string {
	cookie := s.cookie.Cookie(config.Name)
	if cookie == nil {
		var sid string
		sid = s.CreateSessionID()
		s.cookie.SetCookieByExpire(config.Name, sid, config.Expire)
		return sid
	}
	return cookie.String()
}

//Start 启动session
func (s *Session) Start() bool {
	sid := s.SessionID()
	if sid != "" {
		return true
	}
	return false
}

//Name
func (s *Session) Name(key string) string {
	sid := s.SessionID()
	return fmt.Sprintf("%s-%s%s", sid, config.Prefix, key)
}

//Sessioner
type Sessioner interface {
	Set(key string, value interface{}) bool
	Get(key string, value interface{}) bool
	Del(key string) bool
	FlushDB() bool
	Start() bool
	SessionID() string
	SessionKey() string
}

//Cookier
type Cookier interface {
	SetCookies(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
	SetCookie(name, value string)
	Cookie(name string) *lib.Data
	SetCookieByExpire(name, value string, expire int)
}

var initSession Session

var onceDo sync.Once

var config SessionConfig

//init
func init() {
	sessionFile := "config/session.config"
	cfgFile := file.GetContent(sessionFile)
	if cfgFile == "" {
		sessionConfig := SessionConfig{
			Name:   "GOSESSID",
			Type:   "file",
			Expire: 3600,
			Prefix: "",
			Path:   "session",
		}
		file.PutContent(sessionFile, fmt.Sprintf("%v", sessionConfig))
		panic("please setting session config in config/session.config file")
	}
	lib.StringToObject(cfgFile, &config)
}

//GetSession
func GetSession(cookie Cookier) Sessioner {
	onceDo.Do(func() {
		initSession = Session{
			cache:  createCache(),
			cookie: cookie,
		}
	})
	initSession.cookie = cookie
	return &initSession
}

//createCache
func createCache() icache.ICacher {
	switch config.Type {
	case "file":
		return createFileCache()
	case "redis":
		return cache.NewCache("redis")
	case "memcache":
		return cache.NewCache("memcache")
	default:
		return nil
	}
}

//createFileCache
func createFileCache() icache.ICacher {
	path := "runtime/cache/session/"
	if config.Path == "" {
		path = config.Path
	}
	return filecache.NewFileCacheByPath(path)
}
