package interaction

import "fmt"

func SayHello() {
	fmt.Println("Hello There!");
}

func ReceiveFeedback(res string) {
	fmt.Println("Your opinion: ", res);
	fmt.Println("Thanks for your feedback we'll fix it as soon as possible! 🙌");
}