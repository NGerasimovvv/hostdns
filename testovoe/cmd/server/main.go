package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"

	"google.golang.org/grpc"
	pb "test/pb"
)

type server struct {
	pb.UnimplementedHostDnsServiceServer
}

func (s *server) SetHostname(ctx context.Context, req *pb.SetHostnameRequest) (*pb.SetHostnameResponse, error) {
	cmd := exec.Command("hostnamectl", "set-hostname", req.Hostname)
	err := cmd.Run()
	if err != nil {
		return &pb.SetHostnameResponse{Status: "Failed"}, err
	}
	fmt.Println("woooooow")
	return &pb.SetHostnameResponse{Status: "Success"}, nil
}

func (s *server) GetDnsServers(ctx context.Context, req *pb.GetDnsServersRequest) (*pb.GetDnsServersResponse, error) {
	content, err := os.ReadFile("/etc/resolv.conf")
	if err != nil {
		return nil, err
	}

	var dnsServers []string
	for _, line := range strings.Split(string(content), "\n") {
		if strings.HasPrefix(line, "nameserver") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				dnsServers = append(dnsServers, parts[1])
			}
		}
	}

	return &pb.GetDnsServersResponse{DnsServers: dnsServers}, nil
}

func (s *server) AddDnsServer(ctx context.Context, req *pb.AddDnsServerRequest) (*pb.AddDnsServerResponse, error) {
	file, err := os.OpenFile("/etc/resolv.conf", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return &pb.AddDnsServerResponse{Status: "Failed"}, err
	}
	defer file.Close()

	if _, err := file.WriteString("nameserver " + req.DnsServer + "\n"); err != nil {
		return &pb.AddDnsServerResponse{Status: "Failed"}, err
	}

	return &pb.AddDnsServerResponse{Status: "Success"}, nil
}

func (s *server) RemoveDnsServer(ctx context.Context, req *pb.RemoveDnsServerRequest) (*pb.RemoveDnsServerResponse, error) {
	content, err := os.ReadFile("/etc/resolv.conf")
	if err != nil {
		return &pb.RemoveDnsServerResponse{Status: "Failed"}, err
	}

	var newContent []string
	for _, line := range strings.Split(string(content), "\n") {
		if strings.HasPrefix(line, "nameserver") && strings.Contains(line, req.DnsServer) {
			continue
		}
		newContent = append(newContent, line)
	}

	err = os.WriteFile("/etc/resolv.conf", []byte(strings.Join(newContent, "\n")), 0644)
	if err != nil {
		return &pb.RemoveDnsServerResponse{Status: "Failed"}, err
	}

	return &pb.RemoveDnsServerResponse{Status: "Success"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHostDnsServiceServer(s, &server{})
	log.Println("gRPC server is running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
