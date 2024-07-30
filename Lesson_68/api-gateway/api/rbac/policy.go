package rbac

import (
	pgadapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/pkg/errors"
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

    [policy_effect]
    e = some(where (p.eft == allow))

    [matchers]
    m = r.sub == p.sub && (r.obj == p.obj || (r.obj == p.obj + "*" && p.obj != "/e-commerce/admin/" && p.obj != "/e-commerce/user/"))
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
		{"admin", "/e-commerce/admin/", "*"},
		{"user", "/e-commerce/user/", "*"},
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to add policies")
	}

	err = e.SavePolicy()
	if err != nil {
		return nil, errors.Wrap(err, "failed to save policies")
	}

	return e, nil
}
