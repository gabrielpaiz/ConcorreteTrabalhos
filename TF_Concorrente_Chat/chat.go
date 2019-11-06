// Construido como parte da disciplina de Sistemas Distribuidos
// Semestre 2018/2  -  PUCRS - Escola Politecnica
// Estudantes:  Andre Antonitsch e Rafael Copstein
// Professor: Fernando Dotti  (www.inf.pucrs.br/~fldotti)
// Algoritmo baseado no livro:
// Introduction to Reliable and Secure Distributed Programming
// Gabriel Bonatto Justo e Gabriel Pereira Paiz

package main

import "fmt"
import "os"
import "bufio"
//import "time"

import BEB "./BEB"


func Send (block chan struct{}, module BEB.BestEffortBroadcast_Module, ads []string) {
	reader := bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')

	block <- struct {}{}
	send_message := BEB.BestEffortBroadcast_Req_Message{
		Addresses : ads[1:],
		Message : message}

	module.Req <- send_message
	
}

func Recv (module BEB.BestEffortBroadcast_Module, ads []string){
	recv_message := <- module.Ind
	fmt.Println(recv_message, ": ", recv_message.Message)
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
	 
	 for{
		 
		 go Send(block, beb, users)
		 go Recv(beb, users)
		 <- block
	 }



































	// fmt.Println(addresses)

	// beb := BEB.BestEffortBroadcast_Module{
	// 	Req: make(chan BEB.BestEffortBroadcast_Req_Message),
	// 	Ind: make(chan BEB.BestEffortBroadcast_Ind_Message)}

	// beb.Init(addresses[0])

	// // enviador de broadcasts
	// go func() {

	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	var msg string

	// 	for {
	// 		if scanner.Scan() {
	// 			msg = scanner.Text()
	// 		}
	// 		req := BEB.BestEffortBroadcast_Req_Message{
	// 			Addresses: addresses[1:],
	// 			Message:   msg}
	// 		beb.Req <- req
	// 	}
	// }()

	// // receptor de broadcasts
	// go func() {
	// 	for {

	// 		in := <-beb.Ind
	// 		fmt.Printf("Message from %v: %v\n", in.From, in.Message)

	// 	}
	// }()

	// blq := make(chan int)
	// <-blq
}

/*

go run chat.go 127.0.0.1:5001  127.0.0.1:6001

go run chat.go 127.0.0.1:6001  127.0.0.1:5001

*/
