package main

func main() {
	gRPCServer := NewGRPCServer(":9000")
	go gRPCServer.Run()

	httpServer := NewHttpServer(":9001")
	httpServer.Run()
}
