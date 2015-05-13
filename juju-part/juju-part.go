package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func usage() {
	fmt.Println("part partitions a network between two units")
	fmt.Println("juju part wordpress/0 mysql/1 #Partitions wordpress/0 and mysql/1")
	fmt.Println("The partition occurs on the machine, so colocated units will also be affected")
}

func unitPrivateIp(unit string) (string, error) {
	cmd := []string{"run", "--unit", unit, "unit-get private-address"}
	out, err := exec.Command("juju", cmd...).CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.Trim(string(out), "\n"), nil
}

func Part(ip string) []string {
	cmd := fmt.Sprintf("sudo iptables -A INPUT -s %s -j DROP", ip)
	return strings.Split(cmd, " ")
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		flag.Usage()
		return
	}
	unitA := flag.Args()[0]
	unitB := flag.Args()[1]

	unitBIP, err := unitPrivateIp(unitB)
	if err != nil {
		log.Fatalf("failed to get private ip of unitB: %v", err)
	}
	partArgs := []string{"ssh", unitA}
	partArgs = append(partArgs, Part(unitBIP)...)
	log.Println(partArgs)
	_, err = exec.Command("juju", partArgs...).Output()
	if err != nil {
		log.Fatalf("failed to partition: %v", err)
	}
}
