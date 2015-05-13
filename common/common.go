package common

import (
	"fmt"
	"os/exec"
	"strings"
)

func UnitPrivateIp(unit string) (string, error) {
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

func Heal(ip string) []string {
	cmd := fmt.Sprintf("sudo iptables -D INPUT -s %s -j DROP", ip)
	return strings.Split(cmd, " ")
}

func Fix() []string {
	cmd := "sudo tc qdisc del dev eth0 root"
	return strings.Split(cmd, " ")
}

func Flaky() []string {
	cmd := `sudo tc qdisc add dev eth0 root netem loss "20%" "75%"`
	return strings.Split(cmd, " ")
}

func Slow() []string {
	cmd := "sudo tc qdisc add dev eth0 root netem delay 50ms 10ms distribution normal"
	return strings.Split(cmd, " ")
}

func Show() []string {
	cmd := "sudo tc qdisc show dev eth0"
	return strings.Split(cmd, " ")
}
