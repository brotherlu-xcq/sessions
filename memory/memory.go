package memory

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

func NewStore(options *sessions.Options) sessions.Store {
	return &store{}
}

type store struct {
	sessions map[string]*sessions.Session
	options  sessions.Options
}

func (s *store) Get(ctx *app.RequestContext) (*sessions.Session, error) {

}

func (s *store) StartClean() {
	go func() {
		select {}
	}()
}
