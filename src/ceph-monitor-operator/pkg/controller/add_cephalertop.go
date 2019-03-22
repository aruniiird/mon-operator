package controller

import (
	"ceph-monitor-operator/pkg/controller/cephalertop"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, cephalertop.Add)
}
