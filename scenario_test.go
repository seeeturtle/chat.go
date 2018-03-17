package chatgo_test

import (
	"fmt"

	"github.com/seeeturtle/chatgo"
)

func Example() {
	scenarios := []chatgo.Scenario{
		chatgo.CommonScenario{
			Conditions: []chatgo.Condition{
				func(o *chatgo.Object) bool { return true },
			},
			Behaviors: []chatgo.Behavior{
				func(o *chatgo.Object) { fmt.Println("Behavior 1 from CommonScenario 1") },
				func(o *chatgo.Object) { fmt.Println("Behavior 2 from CommonScenario 1") },
			},
		},
		chatgo.CommonScenario{
			Conditions: []chatgo.Condition{
				func(o *chatgo.Object) bool { return false },
			},
			Behaviors: []chatgo.Behavior{
				func(o *chatgo.Object) { fmt.Println("Behavior 1 from CommonScenario 2") },
				func(o *chatgo.Object) { fmt.Println("Behavior 2 from CommonScenario 2") },
			},
			Else: []chatgo.Behavior{
				func(o *chatgo.Object) { fmt.Println("Else Behavior 1 from CommonScenario 2") },
				func(o *chatgo.Object) { fmt.Println("Else Behavior 2 from CommonScenario 2") },
			},
		},
	}

	for _, scenario := range scenarios {
		chatgo.RunScenario(&scenario, nil)
	}

	// Output:
	// Behavior 1 from CommonScenario 1
	// Else Behavior 1 from CommonScenario 2
	// Else Behavior 2 from CommonScenario 2
}
