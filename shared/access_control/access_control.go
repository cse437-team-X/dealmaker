package access_control

import "github.com/casbin/casbin/v2"

type WorkerInstance struct {
	ConfPath string
	PolicyPath string
	enforcer *casbin.Enforcer
}

func (w WorkerInstance) Init() *WorkerInstance {
	enforcer, err := casbin.NewEnforcer(w.ConfPath, w.PolicyPath)
	if err != nil {
		panic(err)
	}
	w.enforcer = enforcer
	return &w
}
