//Package advisor is a daemon to automatically activate the
//autoscaling on AWS if the fleet spot is unstable.
package main

import (
	"github.com/getninjas/spot-advisor/pkg/logic"
	"github.com/getninjas/spot-advisor/pkg/structure"
)

// main invokes all initial functions
// sleep is invoked when the daemon finishes performing its tasks
func main() {
	structure.GetRegion()
	logic.Start()
}
