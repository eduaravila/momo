package test

import (
	"net"
	"time"
)

/*
WaitForHTTP waits for a HTTP server to be available at the given address.
timeout is 5 seconds.
*/
func WaitFor(address string) bool {
	success := make(chan bool)

	go func() {
		for {
			conn, err := net.DialTimeout("tcp", address, time.Second)

			if err != nil {
				time.Sleep(time.Second)

				continue
			}

			if conn != nil {
				success <- true

				return
			}
		}
	}()

	timeout := time.After(5 * time.Second)
	select {
	case <-timeout:
		return false
	case <-success:
		return true
	}

}
