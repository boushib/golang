package main

import "fmt"

type User struct {
	username string
	email string
	password string
	age int
	gender string
}

type Admin struct {
	User
	roles []string
}

func main() {
	// 1:42:17
	// x := 3 + 4i
	// y := 1 + 2i
	// z := x / y
	// a := complex(5, 4)
	// fmt.Println(real(z))
	// fmt.Println(imag(z))
	// fmt.Println(a)

	// s := "Hello Go"
	// b := []byte(s)

	// fmt.Printf("%v - %T\n", b, b)
	// fmt.Printf("%v - %T", string(b), string(b))

    /*
	const (
		_  = iota // ignore first value
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
	)

	fileSize := 2000000000.0
	fmt.Printf("%.2f GB\n", fileSize/GB)
	numbers := []int {2, 6, 51, 56}
	fmt.Printf("- numbers: %v\n", numbers)
	/*numbers[0] = 100
	fmt.Printf("- updated numbers: %v\n", numbers)
	fmt.Printf("- numbers count: %v\n", len(numbers))
	arrayOfArrays := [][]int { {1, 2, 3}, {4, 5, 6}, {7, 8, 9} }
	fmt.Printf("Array of arrays: %v", arrayOfArrays)
     */

	// arrays
	/*
	a := [...]int { 3, 4, 5 }
	b := &a
	b[0] = 300
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	 */

	// slices
	/*
	a := []int { 1, 2, 3}
	b := []int {4, 5, 6}
	c := append(a, b...)

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)

	d := make([]int, 5, 100)
	fmt.Printf("d: %v - capacity: %v\n", d, cap(d))
	*/
	
	// maps
	/*
	user := map[string]float32 {
		"age": 29,
		"score": 200,
		"grade": 88,
	}

	// otherUser := make(map[string]int)

	fmt.Printf("user: %v\n", user)
	user["level"] = 9
	admin := user
	delete(admin, "age")
	fmt.Printf("user: %v\n", user)
	_, ok := user["name"]
	// fmt.Print(name, ok)
	fmt.Print(ok)
	// fmt.Printf("other user: %v\n", otherUser)
	 */

	// struct
	user := User {
		username: "john",
		email: "john@google.com",
		password: "123456",
		age: 29,
		gender: "male",
	}

	admin := Admin{
		roles: []string {"admin", "superadmin"},
	}

	admin.username = "root"
	admin.email = "admin@myapp.com"
	admin.password = "admin@123"

	fmt.Printf("user: %v\n", user)
	// fmt.Printf("username: %v\n", user.username)
	fmt.Printf("admin: %v\n", admin)
}
