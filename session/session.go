package session

import (
	"github.com/gorilla/sessions"
	"net/http"
	"github.com/satori/go.uuid"
	"net/url"
	"github.com/yurencloud/yugo/config"
)

func GetInstance(writer http.ResponseWriter, request *http.Request) *sessions.Session {
	// 先从请求中获取cookie的session-id
	sessionId := ""
	cookie, err := request.Cookie("session-id")
	// 如果session-id不存在，代表会话不存在
	if err != nil || cookie.Value == "" {
		sessionId = uuid.Must(uuid.NewV4()).String()
		cookie := http.Cookie{
			Name: "session-id",
			Value: url.QueryEscape(sessionId),
			Path: "/", HttpOnly: true,
			MaxAge: 3600,
		}
		http.SetCookie(writer, &cookie)
	}else{
		sessionId = cookie.Value
	}

	store := sessions.NewCookieStore([]byte(config.Get("session.key")))
	session, _ := store.Get(request, sessionId)

	return session
}
