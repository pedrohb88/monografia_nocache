package transport

import (
	"fmt"
	"log"
	"monografia/database"
	"net/http"
	"os"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/network"
)

func Benchmark(next http.Handler) http.Handler {

	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var cpuUsage float64
		var netUsage uint64

		testID := r.Header.Get("x-test")
		reqID := r.Header.Get("x-req")

		before, err := cpu.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}
		time.Sleep(time.Duration(1) * time.Second)
		after, err := cpu.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}
		total := float64(after.Total - before.Total)
		cpuUser := float64(after.User-before.User) / total * 100
		cpuSystem := float64(after.System-before.System) / total * 100

		cpuUsage = cpuUser + cpuSystem

		var startBytes, endBytes uint64

		netStats, err := network.Get()
		if err != nil {
			log.Fatal(err)
		}
		for _, n := range netStats {
			if n.Name == "eth0" {
				startBytes = n.RxBytes
				break
			}
		}

		time.Sleep(time.Second)

		netStats, err = network.Get()
		if err != nil {
			log.Fatal(err)
		}
		for _, n := range netStats {
			if n.Name == "eth0" {
				endBytes = n.RxBytes
				break
			}
		}

		netUsage = endBytes - startBytes

		_, err = db.Exec(`
			INSERT INTO benchmark(test, resource, x, y)
			VALUES (?, ?, ?, ?), (?, ?, ?, ?)
		`,
			testID, "cpu", reqID, cpuUsage,
			testID, "net", reqID, float64(netUsage),
		)
		if err != nil {
			log.Fatal(err)
		}

		next.ServeHTTP(w, r)
	})
}
