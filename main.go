package main

import (
	"gitlab.com/test-nested/services/module-test"
	"gitlab.com/test-nested/services/module-test/api"
)

type Bar struct {
	stringer module.Stringer
}

func main() {
	b := Bar{
		stringer: &api.Foo{
			SomeString: "renovate issue #",
			SomeInt:    28540,
		},
	}

	print("hope this repo will help you to solve" + b.stringer.ToString())
}
