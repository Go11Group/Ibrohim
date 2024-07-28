package rbac

import (
	pgadapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/pkg/errors"

	_ "github.com/lib/pq"
)

func Policy() (*casbin.Enforcer, error) {
	db, err := ConnectDB()
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to the database")
	}
	defer db.Close()

	a, err := pgadapter.NewAdapter(db, "postgres", "casbin_rule")
	if err != nil {
		return nil, errors.Wrap(err, "failed to construct adapter")
	}

	txt := `
	[request_definition]
	r = sub, obj, act

	[policy_definition]
	p = sub, obj, act

	[role_definition]
	g = _, _

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = r.sub == p.sub && (r.obj == p.obj || p.obj == "*") && (r.act == p.act || p.act == "*")
	`

	m, err := model.NewModelFromString(txt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to construct model")
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, errors.Wrap(err, "failed to construct enforcer")
	}

	err = e.LoadPolicy()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load policy")
	}

	_, err = e.AddPolicies([][]string{
		{"admin", "*", "*"},
		{"user", "*", "GET"},
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to add policy")
	}

	err = e.SavePolicy()
	if err != nil {
		return nil, errors.Wrap(err, "failed to save policy")
	}
	e.EnableLog(true)

	return e, nil
}
