package main

import("./MCCSemaforo")

type Barber struct
{
	costumers int
	mutex MCCSemaforo.Semaphore
	costumer MCCSemaforo.Semaphore
	barber MCCSemaforo.Semaphore
}

func new Barber () *Barber
{
	b:= &Barber{
		mutex: MCCSemaforo.NewSemaphore(1)
		costumer: MCCSemaforo.NewSemaphore(0)
		barber: MCCSemaforo.NewSemaphore(0)
		costumers: 0;
		}
		return b
}

func fila_costumer(b *Barber)
{
	b.mutex.wait()
	if(b.costumers == n+1)
	{
		b.mutex.signal()
		//balk()
	}
	b.costumers += 1
	b.mutex.signal()

	b.costumer.signal()
	b.barber.wait()
	//getHairCut()


	b.mutex.wait()
		b.costumers -= 1
	b.mutex.signal()
}

func main()
{

}