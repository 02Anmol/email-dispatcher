package main

import "fmt"

type Recipient struct {
	Name  string
	Email string
}

func main() {

	err := loadRecipient("./sample.csv")
	if err != nil {
		fmt.Println(err)
	}

}
