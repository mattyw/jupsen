package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func usage() {
	fmt.Println("slow makes the network on the specified unit slow")
	fmt.Println("juju slow mysql/1 # Makes eth0 on the mysql/1 unit slow")
}

func Slow() []string {
	cmd := "sudo tc qdisc add dev eth0 root netem delay 50ms 10ms distribution normal"
	return strings.Split(cmd, " ")
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		return
	}
	unitA := flag.Args()[0]

	slowArgs := []string{"ssh", unitA}
	slowArgs = append(slowArgs, Slow()...)
	log.Println(slowArgs)
	_, err := exec.Command("juju", slowArgs...).Output()
	if err != nil {
		log.Fatalf("failed to slow: %v", err)
	}
}
