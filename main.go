package main

func main() {
	go producer()
	go consumer()
	select {}
}
