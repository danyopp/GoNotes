//folder structure
/*
	bin - compiled code
	pkg - stores code from other packages so it doesn't have to recompile if not changed
	src - source code
		-git repo
			-repo
			-repo
			//-repo
			


	Be sure these are added to git:
		go.mod - modules included
		go.sum - crypto hash to tell if code has been changed
*/

//helpful commands
/*
	go version
	go help								//
	go help fmt							//help for format command
	go fmt ./...						//format all files in repo
	go env								//print env vars
	go list -m -versions rsc.io/sampler	// lists available versions of the sampler package
	go get								//goes and gets package
	go get rsc.io/sampler@v1.3.1		// gets a specific version of sampler package
	
	go mod init example.com/hello		// creates new mod
	go run .							//builds -> runs -> deletes build
	go run notes.go						//run
	go build .							//builds executable
	go clean							// cleans ex and object files
	go install							// compile and install packages and dependencies
	go test								// runs test files

	*/

//Env Vars
/*
	GOROOT -> shows where go bin is installed
	GOPATH -> path to go code; Not always used these days
		can also use the [go mod init github.com/danyopp/GoNotes] command

*/

//----------------------------------------------------------------------------------------------
//func Scan(a ...interface{}) (n int, err error)    	//function takes any number of arguments in any type and returns two vars
// can use _ to catch a return var you don't plan on using; avoid compiler error for unused vars


//Package declaration
package main

//import other packages
import (
	"fmt"
	"rsc.io/quote"
	//"strconv"
)

//main
func main() {
	fmt.Println("Hello World")
	callPackage()
	declarationTest()
	dataTypes()
	ControlFlow()
	ArrayFunction()
	SliceFunction()
	structExample()

	//Function examples
	a,_ := exampleFunc("Happy", []int{1, 2, 3}...)
	a,_ = exampleFunc("Happy", 1, 2, 3)
	a,_ = exampleFunc("Happy")
	fmt.Println(a)
	fmt.Println("Math Callback: ",domath(add, 3,6))

}

//Everything is passed by value... pointers for reference
func exampleFunc(s string, x ...int) (bool, string) {
	fmt.Println(s)
	fmt.Println(x)
	return true, "Money"

}

/*
func deferExample(){
	openFile()
	defer closeFile()
	//do stuff with file and it will close when function returns
}
*/

func callPackage() {
	fmt.Println("\nCall Package Func")

	fmt.Println(quote.Glass())
}

func declarationTest() {
	fmt.Println("\nDeclaration Test Func")
	// short declaration operator:
	//cannot be used outside of function
	x := 43
	fmt.Printf("Variable x: value=%v type=%T\n", x, x)
	// var declaration, can be used globally, gets type from assignment
	var y = 33
	fmt.Printf("Variable y: value=%v type=%T\n", y, y)
	//seperate declaration and assignment
	var z int	//set to false/0/0.0/""/nil
	z = 22
	fmt.Printf("Variable z: value=%v type=%T\n", z, z)
}

func dataTypes() {
	// primative datatypes = very basic types (bool, int, string)
	// composite datatypes = array, slice, ...
	var a = `This is a raw string literal which is a type of string`
	var b = "This is a regular string"
	var c int
	//create your own type
	type hotdog int //creates new type with int as a base
	var d hotdog
	d = 11

	//types
	var e bool
	e = true

	var f float64
	f = 53.2355

	/* NUMERIC TYPES
	uint8
	uint16
	uint32
	uint64

	int8
	int16
	int32
	int64

	float32
	float64

	complex64
	complex128

	byte - alias for uint8
	rune - alias for int32; character UTF8/1-4 bytes
	*/

	/* STRING TYPES
	`` raw string literal
	"" regular string
	immutable, but can assign a new val
	string is just held as a slice of bytes
	bs := []byte(b)
	*/

	// conversion (not called casting)
	c = int(d) //hotdog to int
	bs := []byte(b)

	//Constants
	const g = 909
	const (
		h = 23.23
		j float32 = 56.55 
	)

	//Iota
	// represents successive untyped constants
	const (
		jan = iota
		feb
		march
	)

	//Bit shifts... probably too clever
	k := 8
	fmt.Println( k << 3) // 2^3 * 2^3 = 2^6
	k = 8
	fmt.Println( k >> 3) // 2^3 / 2^3 = 2^0 = 1
	const(
		_ = iota //throw away 0
		kb = 1 << (iota * 10)
		mb = 1 << (iota * 10)
		gb = 1 << (iota * 10)
	)
	//good for kb to mb to gb shift ... shift of 10
	fmt.Println(kb, mb, gb)


	fmt.Printf("Variable a: value=%v type=%T\n", a, a)
	fmt.Printf("Variable b: value=%v type=%T\n", b, b)
	fmt.Printf("Variable c: value=%v type=%T\n", c, c)
	fmt.Printf("Variable d: value=%v type=%T\n", d, d)
	fmt.Printf("Variable e: value=%v type=%T\n", e, e)
	fmt.Printf("Variable f: value=%v type=%T\n", f, f)
	fmt.Printf("Variable bs: value=%v type=%T\n", bs, bs)
	fmt.Printf("Variable g: value=%v type=%T\n", g, g)
	fmt.Printf("Variable h: value=%v type=%T\n", h, h)
	fmt.Printf("Variable j: value=%v type=%T\n", j, j)
	fmt.Println(jan, feb, march)

}

func ControlFlow(){
	fmt.Println("\nControl Flow")
	//no while loop

	//classic for loop
	a := 0
	for i:=0; i<100; i++{
		a++
	}
	fmt.Println(a)

	//mod for to fit traditional while
	b := false
	for !b {
		b= !b
	}
	fmt.Println(b)

	//range for loops


	//break and goto
	for c:=0;c<4;c++ {
		if c == 1{
			continue
		}
		if c == 3{
			break
		}
		fmt.Println(c)
	}

	//if statement
	
	if d := 3; d < 3 || d > 10{
		fmt.Printf("IF")
	}else if d==3 && d<5{
		fmt.Println("Else If")
	}else{
		fmt.Println("Else")
	}

	//switch
		//No Default fallthrough; can specify with fallthrough keyword
	switch {	//evaluates on true
	case false: 
		fmt.Println("case 1")
	case (2==4): 
		fmt.Println("case 2")
	case true: 
		fmt.Println("case 3")
		fallthrough
	case (3==5):
		fmt.Println("case 4")
		fallthrough
	default:
		fmt.Println("case 5")
	}

	switch "Value"{
	case "Value", "Value2":
		fmt.Println("Found")
	default:
		fmt.Println("Other")
	}

}

func ArrayFunction(){
	fmt.Println("ARRAYFUNCTION")
	//really just a building block for slices or detailed mem layout
	//probs better to just use slices
	//these are passed by value, WATCH OUT
	var x [5]int
	x[3] = 12
	fmt.Println(x)
	fmt.Println(len(x)) //5
}

func SliceFunction(){
	fmt.Println("SliceFunction")

	//Define a slice, not mem optimal
	// x := type{values} is a composite literal
	x := []int{4,5,6,7}
	fmt.Println(x , x[3])
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// can also make a slice using the builtin [make] function but probs easier to use composite literal
	// can be used if you know what the cap you will need to optimize construction
	var z []int
	z = make([]int, 10, 1000)
	fmt.Println("Zarray", len(z), cap(z))



	//iterate through slice
	for i, v := range x {
		fmt.Println(i, v)	//0 4...
	}

	//get specific values
	fmt.Println(x[2:3]) // 6

	//append
		//append is a special builtin because it needs compiler support to ensure typing
	x = append(x, 8, 9 ,10)
	fmt.Println(x)
	//append a slice to a slice
	y := []int{99, 98, 97}
	x = append(x, y...) //unroll y using the ... operator
	fmt.Println(x)
	//delete from a slice
	x = append(x[:3], x[4:]...) //7 at position 3 removed
	fmt.Println(x)

//	multi-D slice
//	mdslice := [][]string

}

/*
func MapsFunction(){
	//ehhh
}*/
type human interface{
	speak()
}

type person struct{
	first string
	last string
}

type agent struct{
	person
	coverID string
}

//METHOD - attached to agent struct
func (a agent) speak(){
	a.person.first = a.person.first + "altered" 
	fmt.Println("I am a secret agent", a.first, a.last)
}

//Struct
func structExample(){
	p1 := person{
		first: "James",
		last: "Smith", //note trailing comma
	}
	fmt.Println(p1)

	//can nest structs for 
	a1 := agent{
		person: person{
			first: "Bob",
			last: "Loblaw",
		},
		coverID: "maximus",
	}
	a1.speak()
	fmt.Println(a1.first, a1.last, a1.coverID) // Notice that the inner type vars got promoted to outer type
											// can specify person struct if there are name collisions
	//can define and fill a struct on the fly to create an anon struct
	Ipoly(a1)
}

//Interface polymorphism
func Ipoly(h human){
	fmt.Println("human called function")
}


/*
//Get from interface back to a specific type
	switch h.(type) {
	case person: h.(person).first
	}
//anon function:
	func(x int){//code here}(42)
//Func Expression
	f := func(x int){fmt.Println("this is an expression", x)}
	f(x)
//Return a function
	func main() {
		x := bar
		i := x()
	}
	func bar() func() int {
		return func() int {
			return 451}
	}
//Callback - passing a func in to another function
*/
func add(x int, y int) int{
		return x + y
	}
	func domath(a func(x int, y int)int, b int, c int)int{
		return a(b,c)
	}

//*/

