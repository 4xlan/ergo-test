package testapp

import (
	"ergo.services/ergo/gen"
)

func CreatetestApp() gen.ApplicationBehavior {
	return &testApp{}
}

type testApp struct{}

// Load invoked on loading application using method ApplicationLoad of gen.Node interface.
func (app *testApp) Load(node gen.Node, args ...any) (gen.ApplicationSpec, error) {
	return gen.ApplicationSpec{
		Name:        "testapp",
		Description: "description of this application",
		Mode:        gen.ApplicationModeTransient,
		Group: []gen.ApplicationMemberSpec{
			{
				Name:    "testsup",
				Factory: factory_testSup,
			},
		},
	}, nil
}

// Start invoked once the application started
func (app *testApp) Start(mode gen.ApplicationMode) {}

// Terminate invoked once the application stopped
func (app *testApp) Terminate(reason error) {}
