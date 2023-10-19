package golangWeb

import(
	"testing"
	"fmt"
	"net/http"
	"net/http/httptest"
	"io"
)

func HelloHandler( writer http.ResponseWriter, request *http.Request){
	fmt.Fprintln(writer, "Hello World!")
}

//mencoba menggunakan package httptest untuk unit testing khusus web di golang
func TestHttp( t *testing.T){
	request:= httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder:= httptest.NewRecorder()

	HelloHandler(recorder, request)

	response:=  recorder.Result()
	body, _:= io.ReadAll(response.Body)
	bodyString:= string(body)

	fmt.Println(bodyString)
}