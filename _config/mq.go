package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/streadway/amqp"
)

func (resource *Resource) CreateMQ() error {
	fmt.Println("start_queue")
	maxRetry := 0
	for {
		fmt.Println("retry_time:", maxRetry)
		var connectamqp string = os.Getenv("RABBITMQURL")
		conn, err := amqp.Dial(connectamqp)
		if err == nil {
			resource.MQ = conn
			resource.MQError = make(chan *amqp.Error) // Channel for noti when rabbit close connection
			resource.MQ.NotifyClose(resource.MQError) // Function rabbitmq noti when rabbit close and sent errorChannel parameter to function for try to reconnect
			ch, err := resource.MQ.Channel()
			if err != nil {
				fmt.Println("err_ch", err)
			}
			resource.MQCh = ch
			// resource.MQCh.NotifyClose(resource.MQError)
			go resource.MQReconnector() // Call function Reconnect
			log.Println("Connection established!")
			//q.openChannel()
			//q.declareQueue()
			return nil
		}

		time.Sleep(1000 * time.Millisecond) //Delay for connect
		maxRetry++
	}
	return nil
}

func (resource *Resource) MQReconnector() {
	fmt.Println("START TO RECONNECT RABBIT")
	var recon bool = true
	for recon {
		err := <-resource.MQError              // When Rabbit close connection channel will be back here and call Connect
		fmt.Println("TRY TO RECONNECT RABBIT") //	Method to connect again and put Connection to main function to process data again
		fmt.Println(err)
		resource.CreateMQ()
		recon = false
	}
}

func (r *Resource) CloseMQ() {
	r.MQCh.Close()
	r.MQ.Close()
	color.Cyan("Close MQ Successfully")
}
