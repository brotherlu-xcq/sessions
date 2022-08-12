package sessions

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"time"
)

const SessionKey = "github.com/hertz-contrib/sessions"

type Options struct {
}

type Store interface {
	Get(ctx *app.RequestContext) (*Session, error)

	StartClean()
}

type Session interface {
	ID() string
	Get(key interface{}) interface{}
	Set(key interface{}, val interface{})
	Delete(key interface{})
	Clear()
	AddFlash(value interface{}, vars ...string)
	Flashes(vars ...string) []interface{}
	LastActiveTime() time.Time
}

type DefaultSession struct {
	Id          string
	Values      map[interface{}]interface{}
	LastActTime time.Time
}

func (s *DefaultSession) ID() string {
	return s.Id
}

func (s *DefaultSession) Get(key interface{}) interface{} {
	return s.Values[key]
}

func (s *DefaultSession) Set(key interface{}, val interface{}) {
	s.Values[key] = val
}

func (s *DefaultSession) Delete(key interface{}) {
	delete(s.Values, key)
}

func (s *DefaultSession) Clear() {
	for key := range s.Values {
		s.Delete(key)
	}
}

func (s *DefaultSession) AddFlash(value interface{}, vars ...string) {

}

func (s *DefaultSession) Flashes(vars ...string) []interface{} {
	return nil
}

func (s *DefaultSession) LastActiveTime() time.Time {
	return s.LastActTime
}

func Sessions(name string, store Store) app.HandlerFunc {
	store.StartClean()
	return func(c context.Context, ctx *app.RequestContext) {
		s, err := store.Get(ctx)
		if err != nil {
			hlog.Warnf("[session middleware] find session error", err)
		} else {
			ctx.Set(SessionKey, s)
		}
		ctx.Next(c)
	}
}

func Default(ctx *app.RequestContext) Session {
	return ctx.MustGet(SessionKey).(Session)
}
