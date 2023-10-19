package golangWeb

import(
	"testing"
	"net/http"
	"fmt"
)

//coba handler
func TestHandler(t *testing.T){
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer,"Hello World!")
	}

	server:= http.Server{
		Addr:"localhost:8080",
		Handler:handler,
	}

	err:=server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}

//coba serverMux
func TestSM( t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter,  request *http.Request){
		fmt.Fprint(writer, "Hello World")
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter,  request *http.Request){
		fmt.Fprint(writer, "Hi!")
	})

	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "Image")
	})

	mux.HandleFunc("/images/thumbnail", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "Thumbnails")
	})

	server:= http.Server{
		Addr:"localhost:8080",
		Handler:mux,
	}

	err:=server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}

//menangani request(ceritanya:v)
func TestReq(t *testing.T)  {
	var handler http.HandlerFunc= func (writer http.ResponseWriter, request *http.Request)  {
		fmt.Fprintln(writer, request.Method)
		fmt.Fprintln(writer, request.RequestURI)
	}

	server:= http.Server{
		Addr:"localhost:8080",
		Handler:handler,
	}

	err:=server.ListenAndServe()
	if err != nil{
		panic(err)
	}

}