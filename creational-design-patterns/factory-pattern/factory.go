package main

func newHumanFactory(humanType string) ihuman {
	switch humanType {
	case "basketballer":
		return &basketballer{}
	case "footballer":
		return &footballer{}
	default:
		return nil
	}
}
