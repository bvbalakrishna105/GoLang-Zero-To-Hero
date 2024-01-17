package main

import (
	"bufio"
	"encoding/binary"
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	//exampleWriteToFile()
	//exampleReadFromAFile()
	//exampleReadLineByLine()
	//exampleJSONSerialization()
	//exampleJSONDeserialization()
	//exampleCSVWriting()
	//exampleCSVReading()
	//exampleBinarySerialization()
	//exampleBinaryDeserialization()
	//exampleGOBSerialization()
	exampleGOBDeserialization()
	//fmt.Println("Welcome to FILE IO Operations")
}

func exampleGOBDeserialization() {
	file, err := os.Open("person.gob")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var person MyPerson

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Name:%v,Age:%v", person.Name, person.Age)
}

type MyPerson struct {
	Name string
	Age  int
}

func exampleGOBSerialization() {
	person := MyPerson{Name: "Beesetti", Age: 28}
	file, err := os.Create("person.gob")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Gob data written to file")
}

func exampleBinaryDeserialization() {
	file, err := os.Open("point.bin")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var point Point
	err = binary.Read(file, binary.LittleEndian, &point)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("X:%v,Y:%v", point.X, point.Y)
}

type Point struct {
	X, Y float64
}

func exampleBinarySerialization() {
	point := Point{X: 3.14, Y: 2.71}

	file, err := os.Create("point.bin")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	err = binary.Write(file, binary.LittleEndian, &point)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Binary data written to file")
}

func exampleCSVReading() {

	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, record := range records {
		fmt.Println("Name:", record[0], "Age:", record[1])
	}
}

func exampleCSVWriting() {
	file, err := os.Create("data.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	data := [][]string{
		{"Name", "Age"},
		{"Beesetti", "40"},
		{"Venky", "39"},
	}

	for _, record := range data {
		err := writer.Write(record)

		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("CSV data written to file")
}

func exampleJSONDeserialization() {
	file, err := os.Open("person.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var person Person

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&person)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Name: %v, Age:%v", person.Name, person.Age)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func exampleJSONSerialization() {

	person := Person{Name: "Venky", Age: 30}
	file, err := os.Create("person.json")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(person)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("JSON data written to file")

}

func exampleReadLineByLine() {
	file, err := os.Open("Output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func exampleReadFromAFile() {
	data, err := os.ReadFile("Output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File content:", string(data))
}

func exampleWriteToFile() {

	file, err := os.Create("Output.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	data := []byte("Hello, File I/O in GO")
	_, err = file.Write(data)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data is writtent into file")

}
