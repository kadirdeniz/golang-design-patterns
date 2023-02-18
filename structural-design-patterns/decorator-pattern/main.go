package main

func main() {

	userObj := userStruct{}

	jeanObj := jean{user: userObj}

	tshirtObj := tshirt{user: jeanObj}

	skirtObj := skirt{user: tshirtObj}

	skirtObj.wear()
}
