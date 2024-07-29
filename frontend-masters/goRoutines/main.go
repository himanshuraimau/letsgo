package main



import (
	"fmt"
	"time"
)


func printMessage(text string, chan1 chan string){
        for i:=0; i<10; i++{
		fmt.Println(text)
		time.Sleep(400*time.Millisecond)
		}
       chan1 <- "Done"
	}

func main(){
   var chan1 chan string
   go printMessage("Hello", chan1)
   response := <- chan1
   fmt.Println(response)
	//go printMessage("World")
	//time.Sleep(5*time.Second)
	//fmt.Println("Done")

}
