package chatgo

type Scenario interface {
	Do(*Object)
	Next(*Object) (*Scenario, *Object)
}

func RunScenario(scenario *Scenario, o *Object) {
	next := scenario
	input := o
	for {
		if next != nil {
			(*next).Do(o)
			next, input = (*next).Next(input)
		} else {
			return
		}
	}
}

type CommonScenario struct {
	Conditions   []func(*Object) bool
	Behaviors    []func(*Object)
	Else         []func(*Object)
	NextScenario *Scenario
	NextInput    *Object
}

func (scenario CommonScenario) Do(o *Object) {
	for i, condition := range scenario.Conditions {
		if condition(o) {
			scenario.Behaviors[i](o)
			return
		}
	}

	for _, elseBehaviors := range scenario.Else {
		elseBehaviors(o)
	}
}

func (scenario CommonScenario) Next(o *Object) (*Scenario, *Object) {
	return scenario.NextScenario, scenario.NextInput
}
