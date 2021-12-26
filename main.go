package main

type HTTPRequest struct {
	Method string
}

func main() {
	r := HTTPRequest{"POST"}

	switch r.Method {
	case "GET":
		println("GET is called")
	case "POST":
		println("POST is called")
	case "PUT":
		println("PUT is called")
	case "DELETE":
		println("DELETE is called")
	default:
		println("Unhandled method")

	}

}
