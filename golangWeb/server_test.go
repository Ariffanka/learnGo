package golangWeb

import(
	"testing"
	"net/http"
)

//coba serve ke localhost:8001
func TestServer( t *testing.T){
	server := http.Server{
		Addr: "localhost:8001",
	}

	err:= server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}