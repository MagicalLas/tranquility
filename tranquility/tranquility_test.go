package tranquility

import "fmt"

func f(name string) string {
	return name
}

func ExampleLogging() {
	newF := f
	Tranquility(&newF, f, func(args ...interface{}) {
		fmt.Println(args)
	})

	fmt.Println(newF("Las"))

	// Output:
	// [Las]
	// Las
}
