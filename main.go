package main

import "rate-limiting/controller"

func main() {
	controller := controller.Server{}

	// controller.RateLimitingController()
	controller.ChannelGoRoutineController()
}
