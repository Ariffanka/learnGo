package gouroutine

import(
	"testing"
	"fmt"
	"time"
)

func TestChannel(t *testing.T){
	channel:= make(chan string)
	defer close(channel)


	//kirim data ke channel(
	// channel <- "arip"

	// var data string
	// data <- channel

	// data:= <- channel

	// fmt.Println(<-channel))

	go func(){
		time.Sleep( 2 * time.Second)
		channel <- "Muhammad Ariffanka"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	data:= <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}