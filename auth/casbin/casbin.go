package casbin

import (
	"context"
	"errors"
	"github.com/EscAlice/go-demo/auth"
	casbinstd "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type casbinContextKey string

const (
	ModelCasbinKey    casbinContextKey = "Model"
	PolicyCasbinKey   casbinContextKey = "Policy"
	EnforcerCasbinKey casbinContextKey = "Enforcer"
	AuthCasbinKey     casbinContextKey = "AuthCasbin"

	defaultCasbinRBACModel = `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
)

var (
	ErrAuthCasbinCreatorMissing = errors.New("authCasbinCreator is required")
	ErrEnforcerMissing          = errors.New("enforcer is missing")
	ErrCasbinParseFailed        = errors.New("security info fault")
	ErrUnauthorized             = errors.New("unauthorized access")
	ErrDbMissing                = errors.New("db is missing")
)

type Option func(*options)

type options struct {
	authCasbinCreator auth.AuthCasbinCreator
	model             model.Model
	policy            *gormadapter.Adapter
	db                *gorm.DB
	enforcer          *casbinstd.SyncedEnforcer
}

func WithAuthCasbinCreator(authC auth.AuthCasbinCreator) Option {
	return func(o *options) {
		o.authCasbinCreator = authC
	}
}

func WithCasbinModel(model model.Model) Option {
	return func(o *options) {
		o.model = model
	}
}

func WithCasbinPolicy(policy *gormadapter.Adapter) Option {
	return func(o *options) {
		o.policy = policy
	}
}

func WithCasbinDB(db *gorm.DB) Option {
	return func(o *options) {
		o.db = db
	}
}

// loadRbacModel 加载RBAC模型
func loadRbacModel() (model.Model, error) {
	return model.NewModelFromString(defaultCasbinRBACModel)
}

func loadRbacPolicy(db *gorm.DB) (*gormadapter.Adapter, error) {
	a, err := gormadapter.NewAdapterByDB(db)
	return a, err
}

func NewCabin(opts ...Option) *casbinstd.SyncedEnforcer {
	o := &options{
		authCasbinCreator: nil,
	}
	for _, opt := range opts {
		opt(o)
	}
	if o.db == nil {
		return nil
	}
	if o.model == nil {
		o.model, _ = loadRbacModel()
	}
	if o.policy == nil {
		o.policy, _ = loadRbacPolicy(o.db)
	}
	o.enforcer, _ = casbinstd.NewSyncedEnforcer(o.model, o.policy)
	return o.enforcer
}

func CheckCasbin(ctx context.Context, opts ...Option) bool {
	o := &options{
		authCasbinCreator: nil,
	}
	for _, opt := range opts {
		opt(o)
	}
	if o.enforcer == nil {
		return false
	}
	if o.authCasbinCreator == nil {
		return false
	}
	authC := o.authCasbinCreator()
	if err := authC.ParseFromContext(ctx); err != nil {
		return false
	}
	ctx = context.WithValue(ctx, AuthCasbinKey, authC)
	allowed, _ := o.enforcer.Enforce(authC.GetSub(), authC.GetObj(), authC.GetAct())
	if !allowed {
		return false
	}
	return true
}
