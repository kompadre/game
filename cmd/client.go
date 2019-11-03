package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"../proto"
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
	defer conn.Close()
	ctx := context.Background()
	client := proto.NewSessionClient(conn)
	req := proto.SessionRequest{Username: "kompadre", Password: "Unlikely"}
	resp, err := client.NewSession(ctx, &req)
	if resp.Uuid == "" {
		fmt.Printf("We don't have a session! [%s]\n", resp.GetReason())
		return
	}
	uuid := resp.Uuid
	fmt.Printf("Uuid received: %v\n", uuid)
	done := make(chan bool)
	stream, err := client.LookAround(ctx)

	firstreq := proto.LookAroundRequest{Myuuid: uuid, Type: proto.Types_SYSTEM}
	err = stream.Send(&firstreq)
	fmt.Printf("First req sent: %v", firstreq)

	go func() {
		ctx := stream.Context()
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			resp, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("Closing connection")
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v\n", err)
			}

			if resp.Type == proto.Types_CHAT {
				fmt.Printf("Received a chat: \n")
				fmt.Printf("\t%v\n", resp)
			} else {
				/*
					fmt.Printf("Received look around: \n")
					for key := range resp.Results {
						fmt.Printf("\t%v\n", resp.Results[key].Uuid)
					}
				*/
			}
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
