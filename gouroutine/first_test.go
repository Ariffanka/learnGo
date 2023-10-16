package gouroutine

import(
	"fmt"
	"testing"
	"time"
)

func Hello(){
	fmt.Println("Hello brader")
}

func TestGouroutine(t *testing.T){
	go Hello()
	fmt.Println("Ups")

	time.Sleep( 1 * time.Second)
}

func DisplayNumber(number int){
	fmt.Println("Display", number)
}

func TestMany(t *testing.T){
	for i := 0; i < 100000; i++{
		DisplayNumber(i)
	}
}

func TestManyGou(t *testing.T){
	for i := 0; i < 100000; i++{
		go DisplayNumber(i)
	}

	time.Sleep(5* time.Second)
}