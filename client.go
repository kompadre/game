package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"./proto"
	"google.golang.org/grpc"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:8080", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		fail(err)
	}
	//defer conn.Close()

	client := proto.NewSessionClient(conn)
	ctx := context.Background()
	stream, err := client.NewSession(ctx)
	if err != nil {
		fail(err)
	}
	//req := proto.SessionRequest{Username: "kompadre", Password: "Unlikely"}

	/*
		if err := stream.Send(&req); err != nil {
			log.Fatalf("cannot send %v", err)
		}
		resp, err := stream.Recv()
		fmt.Printf("New session received: %v", resp)
		if err := stream.Send(&req); err != nil {
			log.Fatalf("cannot send %v", err)
		}
		resp, err = stream.Recv()
		fmt.Printf("New session received: %v", resp)
	*/

	done := make(chan bool)
	//	send := make(chan bool)
	uuid := ""
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("-> ")
			var newreq *proto.SessionRequest = nil
			if uuid == "" {
				text, _ := reader.ReadString('\n')
				text = strings.Replace(text, "\n", "", -1)
				password := text
				fmt.Printf("Password is %v.\n", password)
				newreq = &proto.SessionRequest{Username: "kompadre", Password: password}
				if err := stream.Send(newreq); err != nil {
					log.Fatalf("cannot send %v\n", err)
				}
			} else {
				newreq := &proto.LookAroundRequest{}
				if err := stream.Send(newreq); err != nil {
					log.Fatalf("cannot send %v\n", err)
				}
			}

		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v\n", err)
			}
			uuid := resp.Uuid
			if uuid == "" {
				log.Fatalf("No UUID, Error: %v", resp.Reason)
			}
			fmt.Printf("New UUID received: %v\n", uuid)
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done
	log.Println("finished")
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
