package sqlstore

import (
	"context"
	"reflect"

	"xorm.io/xorm"

	"github.com/grafana/grafana/pkg/infra/log"
)

var sessionLogger = log.New("sqlstore.session")

type DBSession struct {
	*xorm.Session
	transactionOpen bool
	events          []interface{}
}

type DBTransactionFunc func(sess *DBSession) error

func (sess *DBSession) publishAfterCommit(msg interface{}) {
	sess.events = append(sess.events, msg)
}

// NewSession returns a new DBSession
func (ss *SQLStore) NewSession(ctx context.Context) *DBSession {
	sess := &DBSession{Session: ss.engine.NewSession()}
	sess.Session = sess.Session.Context(ctx)
	return sess
}

func newSession(ctx context.Context) *DBSession {
	sess := &DBSession{Session: x.NewSession()}
	sess.Session = sess.Session.Context(ctx)

	return sess
}

func startSessionOrUseExisting(ctx context.Context, engine *xorm.Engine, beginTran bool) (*DBSession, bool, error) {
	value := ctx.Value(ContextSessionKey{})
	var sess *DBSession
	sess, ok := value.(*DBSession)

	if ok {
		sessionLogger.Debug("reusing existing session", "transaction". sess.transactionOpen)
		sess.Session = sess.Session.Context(ctx)
		return sess, false, nil
	}

	newSess := &DBSession{Session: engine.NewSession(), transactionOpen: beginTran}
	if beginTran {
		err := newSess.Begin()
		if err != nil {
			return nil, false, err
		}
	}

	newSess.Session = newSess.Session.Context(ctx)
	return newSess, true, nil
}

func getTypeName(bean interface{}) (res string) {
	t := reflect.TypeOf(bean)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}
