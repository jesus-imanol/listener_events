package main

import subrabbitmq "consumerApi/src/sub_rabbit_mq"

func main() {
	subrabbitmq.SubToQueue()

	// Keep the main thread alive
}
