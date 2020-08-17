package main

import (
	"fmt"
	"log"
	"org.quickstart.etcd/example4/discovery"
	//dis "org.quickstart.etcd/example4/discovery"
	"time"
)

func main() {
	//testMaster()
	testService()
}

func testService() {

	serviceName := "s-test"
	serviceInfo := discovery.ServiceInfo{IP: "192.168.1.26"}

	s, err := discovery.NewService(serviceName, serviceInfo, []string{
		"http://192.168.1.17:2379",
		"http://192.168.1.17:2479",
		"http://192.168.1.17:2579",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name:%s, ip:%s\n", s.Name, s.Info.IP)

	go func() {
		time.Sleep(time.Second * 20)
		s.Stop()
	}()

	s.Start()
}
