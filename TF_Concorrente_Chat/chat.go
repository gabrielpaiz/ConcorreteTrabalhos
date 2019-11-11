
// Gabriel Bonatto Justo e Gabriel Pereira Paiz

package main

import "fmt"
import "os"
import "bufio"
import "time"

import BEB "./BEB"


func Send (block chan struct{}, module BEB.BestEffortBroadcast_Module, ads []string) {
	for{
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')

		send_message := BEB.BestEffortBroadcast_Req_Message{
			Addresses : ads[1:],
			Message : message}

		module.Req <- send_message
	}
	
	
}

func Recv (module BEB.BestEffortBroadcast_Module, ads []string){
	for{
		recv_message := <- module.Ind
		fmt.Println(recv_message.From, ": ", recv_message.Message)
	}
	
}






func join (module BEB.BestEffortBroadcast_Module, ads []string){
	message := ads[0] + " entrou no chat"

	send_message := BEB.BestEffortBroadcast_Req_Message{
		Addresses : ads[1:],
		Message : message}

	module.Req <- send_message
}

func main() {


	 if len(os.Args) < 2 {
	 	fmt.Println("Please specify at least two address:port!")
	 	return
	 }

	 users := os.Args[1:]
	 fmt.Println("users : ", users)


	 beb := BEB.BestEffortBroadcast_Module {
		Ind : make(chan BEB.BestEffortBroadcast_Ind_Message),
		Req : make(chan BEB.BestEffortBroadcast_Req_Message)}

	 beb.Init(users[0])

	 block := make(chan struct{})

	 go join(beb, users)
	 
		 
	go Send(block, beb, users)
	go Recv(beb, users)

	for{
		time.Sleep(2 * time.Second)
	}


}

/*

go run chat.go 127.0.0.1:5001  127.0.0.1:6001

go run chat.go 127.0.0.1:6001  127.0.0.1:5001

*/
