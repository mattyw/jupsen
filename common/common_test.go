package common_test

import (
	"testing"

	"github.com/mattyw/jupsen/common"
)

func compareArgs(v, expect []string) bool {
	if len(v) != len(expect) {
		return false
	}
	for i, x := range v {
		if expect[i] != x {
			return false
		}
	}
	return true
}

func TestPart(t *testing.T) {
	args := common.Part("10.0.0.1")
	expect := []string{"sudo", "iptables", "-A", "INPUT", "-s", "10.0.0.1", "-j", "DROP"}
	if !compareArgs(args, expect) {
		t.Errorf("Expected %s got %s", expect, args)
	}
}

func TestHeal(t *testing.T) {
	args := common.Heal("10.0.0.1")
	expect := []string{"sudo", "iptables", "-D", "INPUT", "-s", "10.0.0.1", "-j", "DROP"}
	if !compareArgs(args, expect) {
		t.Errorf("Expected %s got %s", expect, args)
	}
}

func TestFix(t *testing.T) {
	args := common.Fix()
	expect := []string{"sudo", "tc", "qdisc", "del", "dev", "eth0", "root"}
	if !compareArgs(args, expect) {
		t.Errorf("Expected %s got %s", expect, args)
	}
}

func TestFlaky(t *testing.T) {
	args := common.Flaky()
	expect := []string{"sudo", "tc", "qdisc", "add", "dev", "eth0", "root", "netem", "loss", `"20%"`, `"75%"`}
	if !compareArgs(args, expect) {
		t.Errorf("Expected %s got %s", expect, args)
	}
}

func TestSlow(t *testing.T) {
	args := common.Slow()
	expect := []string{"sudo", "tc", "qdisc", "add", "dev", "eth0", "root", "netem", "delay", "50ms", "10ms", "distribution", "normal"}
	if !compareArgs(args, expect) {
		t.Errorf("Expected %s got %s", expect, args)
	}
}

func TestShow(t *testing.T) {
	args := common.Show()
	expect := []string{"sudo", "tc", "qdisc", "show", "dev", "eth0"}
	if !compareArgs(args, expect) {
		t.Errorf("Expected %s got %s", expect, args)
	}
}
