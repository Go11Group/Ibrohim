package rbac

import (
	pgadapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
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

	e, err := casbin.NewEnforcer("../../config/model.conf", a)
	if err != nil {
		return nil, errors.Wrap(err, "failed to construct enforcer")
	}

	err = e.LoadPolicy()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load policy")
	}

	_, err = e.AddPolicies([][]string{
		{"admin", "/e-commerce/admin/*", "*"},
		{"user", "/e-commerce/user/*", "*"},
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to add policy")
	}

	err = e.SavePolicy()
	if err != nil {
		return nil, errors.Wrap(err, "failed to save policy")
	}

	return e, nil
}
