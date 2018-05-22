package chatgo_test

import (
	"fmt"

	"github.com/seeeturtle/chatgo"
)

func ExampleRunScenario_CommonScenario() {
	scenario1 := chatgo.CondScenario{}
	scenario1.Add(
		func(o chatgo.Object) bool { return true },
		func(o chatgo.Object) (chatgo.Scenario, chatgo.Object) {
			fmt.Println("Behavior 1 from CommonScenario 1")
			return nil, nil
		},
	)
	scenario1.Add(
		func(o chatgo.Object) bool { return true },
		func(o chatgo.Object) (chatgo.Scenario, chatgo.Object) {
			fmt.Println("Behavior 2 from CommonScenario 1")
			return nil, nil
		},
	)

	chatgo.RunScenario(&scenario1, nil)

	// Output:
	// Behavior 1 from CommonScenario 1
}

func ExampleRunScenario_Else() {
	scenario2 := chatgo.CondScenario{}
	scenario2.Add(
		func(o chatgo.Object) bool { return false },
		func(o chatgo.Object) (chatgo.Scenario, chatgo.Object) {
			fmt.Println("Behavior 1 from CommonScenario 2")
			return nil, nil
		},
	)
	scenario2.Else(
		func(o chatgo.Object) (chatgo.Scenario, chatgo.Object) {
			fmt.Println("Else Behavior 1 from CommonScenario 2")
			return nil, nil
		},
	)

	chatgo.RunScenario(&scenario2, nil)

	// Output:
	// Else Behavior 1 from CommonScenario 2
}

func ExmapleRunScenario_Deep() {
	var scenario3, scenario4 chatgo.CondScenario
	scenario3.Add(
		func(o chatgo.Object) bool { return true },
		func(o chatgo.Object) (chatgo.Scenario, chatgo.Object) {
			fmt.Println("Behavior 1 from CommonScenario 3")
			return scenario4, nil
		},
	)

	scenario4.Add(
		func(o chatgo.Object) bool { return true },
		func(o chatgo.Object) (chatgo.Scenario, chatgo.Object) {
			fmt.Println("Behavior 1 from CommonScenario 4")
			return nil, nil
		},
	)

	chatgo.RunScenario(&scenario3, nil)

	// Output:
	// Behavior 1 from CommonScenario 3
	// Behavior 1 from CommonScenario 4
}
