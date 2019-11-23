// Gabriel Bonatto Justo e Gabriel Pereira Paiz

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	BEB "./BEB"
)

//Send from BEB
func Send(block chan struct{}, module BEB.BestEffortBroadcast_Module, ads []string) {
	for {

		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')

		sendMessage := BEB.BestEffortBroadcast_Req_Message{
			Addresses: ads[1:],
			Message:   message}

		module.Req <- sendMessage
	}

}

//Recv from BEB
func Recv(module BEB.BestEffortBroadcast_Module, ads []string) {
	for {
		recvMessage := <-module.Ind
		aux := false
		for i := 1; i < len(ads); i++ {
			if ads[i] == recvMessage.Message[2:] {
				aux = true
			}
		}

		if recvMessage.Message[0:2] == "PP" && ads[1] != recvMessage.Message[2:] && !aux {
			fmt.Println("Teste")
			ads = append(ads, recvMessage.Message[2:])
			sendPara := BEB.BestEffortBroadcast_Req_Message{
				Addresses: ads[1:],
				Message:   "Adicionei 1\n"}
			module.Req <- sendPara
		}

		fmt.Println(recvMessage.From, ": ", recvMessage.Message)
	}

}

func join(module BEB.BestEffortBroadcast_Module, ads []string) {
	message := ads[0] + " entrou no chat"

	perm := "PP" + ads[0]

	sendPermission := BEB.BestEffortBroadcast_Req_Message{
		Addresses: ads[1:2],
		Message:   perm}
	module.Req <- sendPermission

	sendMessage := BEB.BestEffortBroadcast_Req_Message{
		Addresses: ads[1:],
		Message:   message}

	module.Req <- sendMessage
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please specify at least two address:port!")
		return
	}

	users := os.Args[1:]
	fmt.Println("users : ", users)

	beb := BEB.BestEffortBroadcast_Module{
		Ind: make(chan BEB.BestEffortBroadcast_Ind_Message),
		Req: make(chan BEB.BestEffortBroadcast_Req_Message)}

	beb.Init(users[0])

	block := make(chan struct{})

	go join(beb, users)

	go Send(block, beb, users)
	go Recv(beb, users)

	for {
		time.Sleep(2 * time.Second)
	}

}

/*

go run chat.go 127.0.0.1:5001  127.0.0.1:6001

go run chat.go 127.0.0.1:6001  127.0.0.1:5001

go run chat.go 127.0.0.1:4001  127.0.0.1:6001

*/
