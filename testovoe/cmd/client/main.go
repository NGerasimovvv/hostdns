package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "test/pb"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "cli",
		Short: "CLI client for interacting with gRPC server",
	}

	var setHostnameCmd = &cobra.Command{
		Use:   "set-hostname [hostname]",
		Short: "Set the hostname of the system",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			hostname := args[0]

			conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()

			client := pb.NewHostDnsServiceClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			res, err := client.SetHostname(ctx, &pb.SetHostnameRequest{Hostname: hostname})
			if err != nil {
				log.Fatalf("could not set hostname: %v", err)
			}
			fmt.Printf("Response from server: %s\n", res.Status)
		},
	}

	var getDnsServersCmd = &cobra.Command{
		Use:   "get-dns-servers",
		Short: "Get list of DNS servers",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				fmt.Printf("Failed to connect to gRPC server: %v\n", err)
				os.Exit(1)
			}
			defer conn.Close()
			client := pb.NewHostDnsServiceClient(conn)

			res, err := client.GetDnsServers(context.Background(), &pb.GetDnsServersRequest{})
			if err != nil {
				fmt.Printf("Error getting DNS servers: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("DNS Servers:", res.DnsServers)
		},
	}

	var addDnsServerCmd = &cobra.Command{
		Use:   "add-dns-server [dns_server]",
		Short: "Add a DNS server",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dnsServer := args[0]
			conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				fmt.Printf("Failed to connect to gRPC server: %v\n", err)
				os.Exit(1)
			}
			defer conn.Close()
			client := pb.NewHostDnsServiceClient(conn)

			res, err := client.AddDnsServer(context.Background(), &pb.AddDnsServerRequest{DnsServer: dnsServer})
			if err != nil {
				fmt.Printf("Error adding DNS server: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("DNS server added:", res.Status)
		},
	}

	var removeDnsServerCmd = &cobra.Command{
		Use:   "remove-dns-server [dns_server]",
		Short: "Remove a DNS server",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			dnsServer := args[0]
			conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				fmt.Printf("Failed to connect to gRPC server: %v\n", err)
				os.Exit(1)
			}
			defer conn.Close()
			client := pb.NewHostDnsServiceClient(conn)

			res, err := client.RemoveDnsServer(context.Background(), &pb.RemoveDnsServerRequest{DnsServer: dnsServer})
			if err != nil {
				fmt.Printf("Error removing DNS server: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("DNS server removed:", res.Status)
		},
	}

	rootCmd.AddCommand(setHostnameCmd)
	rootCmd.AddCommand(getDnsServersCmd)
	rootCmd.AddCommand(addDnsServerCmd)
	rootCmd.AddCommand(removeDnsServerCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
