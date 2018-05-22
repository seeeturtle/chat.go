package chatgo

type Condition func(Object) bool
type Behavior func(Object) (Scenario, Object)

type Scenario interface {
	Next(Object) (Scenario, Object)
}

func RunScenario(scenario Scenario, o Object) Object {
	next := scenario
	input := o
	for {
		if next != nil {
			next, input = next.Next(input)
		} else {
			return input
		}
	}
}

type CondScenario struct {
	conditions   []Condition
	behaviors    []Behavior
	elseBehavior Behavior
}

func (scenario *CondScenario) Add(condition Condition, behavior Behavior) {
	scenario.conditions = append(scenario.conditions, condition)
	scenario.behaviors = append(scenario.behaviors, behavior)
}

func (scenario *CondScenario) Else(behavior Behavior) {
	scenario.elseBehavior = behavior
}

func (scenario CondScenario) Next(o Object) (Scenario, Object) {
	for i, condition := range scenario.conditions {
		if condition(o) {
			return scenario.behaviors[i](o)
		}
	}

	if scenario.elseBehavior != nil {
		return scenario.elseBehavior(o)
	} else {
		return nil, nil
	}
}
