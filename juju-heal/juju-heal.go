package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func usage() {
	fmt.Println("heal removes partitions between two units")
	fmt.Println("juju heal wordpress/0 mysql/1 #Heals partitions between wordpress/0 and mysql/1")
}

func unitPrivateIp(unit string) (string, error) {
	cmd := []string{"run", "--unit", unit, "unit-get private-address"}
	out, err := exec.Command("juju", cmd...).CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.Trim(string(out), "\n"), nil
}

func Heal(ip string) []string {
	cmd := fmt.Sprintf("sudo iptables -D INPUT -s %s -j DROP", ip)
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
	healArgs := []string{"ssh", unitA}
	healArgs = append(healArgs, Heal(unitBIP)...)
	log.Println(healArgs)
	_, err = exec.Command("juju", healArgs...).Output()
	if err != nil {
		log.Fatalf("failed to heal: %v", err)
	}
}
