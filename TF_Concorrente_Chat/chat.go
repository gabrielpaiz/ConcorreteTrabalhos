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
//Função que Faz o Envio de Mensagem
func Send(module BEB.BestEffortBroadcast_Module, ads *[]string, messages *[]string) {
	for {

		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')

		sendMessage := BEB.BestEffortBroadcast_Req_Message{
			Addresses: (*ads)[1:],
			Message:   message}

		module.Req <- sendMessage
		*messages = append((*messages), ((*ads)[0] + ":"+ sendMessage.Message))
	}

}

//Recv from BEB
//Função que Faz Recebimento da Mensagem
func Recv(module BEB.BestEffortBroadcast_Module, ads *[]string, messages *[]string) {
	for {
		recvMessage := <-module.Ind
		aux := true

		//Verifica se o endereço já existe no vetor ads
		for i := 1; i < len(*ads); i++ {
			//fmt.Println("mesage[1:]: ", recvMessage.Message[2:])
			//fmt.Println("ads[i]: ", (*ads)[i])
			if (*ads)[i] == recvMessage.Message[2:] {
				aux = false
			}
		}

		// Verifica se a menssagem recebida é de um novo usuario
		if recvMessage.Message[0:2] == "PP" && aux {
			//fmt.Println("Teste")
			*ads = append((*ads), recvMessage.Message[2:])

			sendPara := BEB.BestEffortBroadcast_Req_Message{
				Addresses: (*ads)[1:],
				Message:   recvMessage.Message}
			module.Req <- sendPara

			fmt.Println("len: ", len(*messages))
			for i := 0; i < len(*messages); i++{
				fmt.Println("entri no for do historico")
				sendHist := BEB.BestEffortBroadcast_Req_Message{
					Addresses: (*ads)[len((*ads))-1 : len((*ads))],
					Message: (*messages)[i]}

				module.Req <- sendHist

				//time.Sleep(100 * time.Millisecond)
				//for i:=0; i<500; i++{}
			}
		}else{
			*messages = append(*messages, (recvMessage.From + ":" + recvMessage.Message))
			fmt.Println(recvMessage.From, ": ", recvMessage.Message)
		}
	}

}

// Funcção de entrada
func join(module BEB.BestEffortBroadcast_Module, ads []string) {
	message := "\n" + ads[0] + " entrou no chat\n"

	perm := "PP" + ads[0]
	sendMessage := BEB.BestEffortBroadcast_Req_Message{
		Addresses: ads[1:],
		Message:   message}

	module.Req <- sendMessage

	
	for i:=0; i<500; i++{} //Evita que as duas mensagens sejam enviadas juntas (com sleep n~ao funciona)

	sendPermission := BEB.BestEffortBroadcast_Req_Message{
		Addresses: ads[1:2],
		Message:   perm}
	module.Req <- sendPermission


	


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


	join(beb, users)

	var messages [] string
	go Send(beb, &users, &messages)
	go Recv(beb, &users, &messages)

	for {
		time.Sleep(2 * time.Second)
	}

}

/*

go run chat.go 127.0.0.1:5001  127.0.0.1:6001

go run chat.go 127.0.0.1:6001  127.0.0.1:5001

go run chat.go 127.0.0.1:4001  127.0.0.1:6001

*/
