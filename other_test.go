package godence

import "testing"

func TestMain(m *testing.M) {
	// start: Init a flow client.
	initFlowClient()
	m.Run()
}
