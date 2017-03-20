// Copyright (c) 2016 Magnus Bäck <magnus@noun.se>

package logstash

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

// If a TestMain function is present, this is executed instead of running all the tests,
// if `go test` is executed. This allows us to use the test binary as a logstashMock
// in TestParallelProcess. The logstashMock is executed, if the env var
// `TEST_MAIN=logstash-mock` is set.
func TestMain(m *testing.M) {
	switch os.Getenv("TEST_MAIN") {
	case "logstash-mock":
		logstashMock()
	default:
		os.Exit(m.Run())
	}
}

// lostashMock returns all input (unprocessed), received on the unix domain socket provided in
// the env var `TEST_SOCKET` via stdout.
func logstashMock() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)
	go func() {
		for _ = range c {
			fmt.Println("Ignoring HUP Signal!")
		}
	}()

	ready := false
	for !ready {
		if _, err := os.Stat(os.Getenv("TEST_SOCKET")); err == nil {
			ready = true
		}
	}

	// TODO: Improve logstash-mock to support control socket
	// conn, err := net.Dial("unix", os.Getenv("CONTROL_SOCKET"))
	// if err != nil {
	// 	log.Fatalf("Failed to dial %s with error: %s", os.Getenv("CONTROL_SOCKET"), err)
	// }

	conn, err := net.Dial("unix", os.Getenv("TEST_SOCKET"))
	if err != nil {
		log.Fatalf("Failed to dial %s with error: %s", os.Getenv("CONTROL_SOCKET"), err)
	}

	b, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalf("Eror while reading from socket: %s", err)
	}
	fmt.Print(string(b))
}
