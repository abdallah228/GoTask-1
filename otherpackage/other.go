package otherpackage

import "fmt"

var Welcome string

func OtherPackage() {
	Welcome = "task 5 welcoming msg in other Package"
	fmt.Println(Welcome)

}
