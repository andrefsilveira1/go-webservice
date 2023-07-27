package main

func main() {
	service := NewCatFactService("https://catfact.ninja/fact")
	service = NewLoggerService(service)
}
