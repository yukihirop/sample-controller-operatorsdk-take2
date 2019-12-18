package controller

import (
	"github.com/yukihirop/sample-controller-operatorsdk-take2/pkg/controller/foo"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, foo.Add)
}
