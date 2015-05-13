package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func usage() {
	fmt.Println("fix fixes any network problems jupsen has created")
	fmt.Println("juju fix mysql/1 # Fixes any jupsen issues created on mysql/1")
}

func Fix() []string {
	cmd := "sudo tc qdisc del dev eth0 root"
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

	fixArgs := []string{"ssh", unitA}
	fixArgs = append(fixArgs, Fix()...)
	log.Println(fixArgs)
	_, err := exec.Command("juju", fixArgs...).Output()
	if err != nil {
		log.Fatalf("failed to fix: %v", err)
	}
}
