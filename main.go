package main

// import "fmt"

func main()  {
	// fmt.Println("main go")
	server := NewServer("127.0.0.1",8080)
	server.Start()
}