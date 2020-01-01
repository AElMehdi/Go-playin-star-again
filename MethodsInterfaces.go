package main

import "fmt"

func main() {
	methods()
}

type MyInt int

type User struct {
	name string
	age int
}

func (u User) isAdult() bool {
	return u.age >= 18
}

func (age MyInt) isAdult() bool {
	return age >= 18
}

func isAdult(user User) bool {
	return user.age >= 18
}

func methods() {
	anAdult := User{"Dean", 45,}
	fmt.Println(anAdult.isAdult())
	fmt.Println(isAdult(anAdult))
	
	anAge := MyInt(13)
	fmt.Println(anAge.isAdult())
}
