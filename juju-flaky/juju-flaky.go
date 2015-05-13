package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func usage() {
	fmt.Println("flaky drops packets on the specified unit")
	fmt.Println("juju flaky mysql/1 # Drops packets on mysql/1")
}

func Flaky() []string {
	cmd := `sudo tc qdisc add dev eth0 root netem loss "20%" "75%"`
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

	flakyArgs := []string{"ssh", unitA}
	flakyArgs = append(flakyArgs, Flaky()...)
	log.Println(flakyArgs)
	_, err := exec.Command("juju", flakyArgs...).Output()
	if err != nil {
		log.Fatalf("failed to fix: %v", err)
	}
}
