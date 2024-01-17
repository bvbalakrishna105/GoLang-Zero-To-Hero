package main

import "fmt"

func main() {

	//exampleBasicInterface()
	//exampleMultiInterface()
	//examplePloyWithInterface()
	//exampleCustomStringWithInterface()
	//exampleTypeAssertionAndSwitch()
	//exampleInterfaceWithPointer()
	//exampleCustomErrorInterface()
	//exampleInterfaceEmbedding()
	//exampleInterfaceConversion()
	exampleTypeAssertionWithPoiterInterface()
	//fmt.Println("Welcome to Interface and Polymorphism")
}

// ==================================================================================
type MyShape interface {
	Area() float64
}

type MyCircle struct {
	Radius float64
}

func (c MyCircle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func exampleTypeAssertionWithPoiterInterface() {
	var s MyShape
	s = &MyCircle{Radius: 5}

	if circle, ok := s.(*MyCircle); ok {
		fmt.Println("Circle Area", circle.Area())
	}
}

// ==================================================================
type Speaker interface {
	Speak()
}

type MyDog struct{}

func (d MyDog) Speak() {
	fmt.Println("Woof!!")
}

type MyCat struct{}

func (c MyCat) Speak() {
	fmt.Println("Meow!!")
}

func exampleInterfaceConversion() {
	animals := []interface{}{MyDog{}, MyCat{}}
	for _, animal := range animals {
		if speaker, ok := animal.(Speaker); ok {
			speaker.Speak()
		}
	}
}

// =======================================================================
type MyReader interface {
	Read() string
}

type MyWriter interface {
	Write(string) error
}

type ReadWriter interface {
	MyReader
	MyWriter
}

type MyReadWrite struct{}

func (rw MyReadWrite) Read() string {
	return "Reading"
}

func (rw MyReadWrite) Write(msg string) error {
	fmt.Println("Writing ", msg)
	return nil
}

func exampleInterfaceEmbedding() {
	var rw ReadWriter = MyReadWrite{}
	fmt.Println(rw.Read())
	fmt.Println(rw.Write("Beesetti Venkata Balakrushna"))
}

// =========================================================================
type MyError struct {
	Msg string
}

func (e MyError) Error() string {
	return e.Msg
}

func Process() error {
	return MyError{Msg: " An Error Occurred"}
}

func exampleCustomErrorInterface() {
	err := Process()

	fmt.Println("Error :", err)
}

// =====================================================================
type Messenger interface {
	GetMessage() string
}

type Email struct {
	Body string
}

func (e *Email) GetMessage() string {
	return e.Body
}

func exampleInterfaceWithPointer() {
	var m Messenger
	email := Email{Body: "Hello here"}

	m = &email

	fmt.Println(m.GetMessage())
}

// ============================================================
func PrintType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Type : int, Value:", v)
	case string:
		fmt.Println("Type : string, Value:", v)

	default:
		fmt.Println("Unknown Type")
	}
}

func exampleTypeAssertionAndSwitch() {
	PrintType(42)
	PrintType("Beesetti")
}

// =============================================================================
type Person struct {
	FirstName, LastName string
}

func (p Person) string() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func exampleCustomStringWithInterface() {
	person := Person{FirstName: "Beesetti", LastName: "Venky"}
	fmt.Println(person)
}

// ========================================================================
type Animal interface {
	MakeSound() string
}

type Dog struct{}
type Cat struct{}

func (d Dog) MakeSound() string {
	return "Woof!!!"
}

func (c Cat) MakeSound() string {
	return "Meow!!"
}

func examplePloyWithInterface() {
	animals := []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		fmt.Println(animal.MakeSound())
	}
}

// ======================================================================
type Writer interface {
	Write(string) error
}

type Logger interface {
	Log(string)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(msg string) error {
	fmt.Println(msg)
	return nil
}

func (cw ConsoleWriter) Log(msg string) {
	fmt.Println("Logging:", msg)
}

func exampleMultiInterface() {
	var w Writer
	w = ConsoleWriter{}
	w.Write("Hello World")

	var l Logger
	l = ConsoleWriter{}
	l.Log("Log message")

}

// ===========================================================================
type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func exampleBasicInterface() {
	var s Shape
	s = Circle{Radius: 5}

	fmt.Println("Circle Area :", s.Area())
}
