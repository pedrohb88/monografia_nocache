package transport

import (
	"context"
	"fmt"
	"log"
	"monografia/database"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/network"
)

func Benchmark(next http.Handler) http.Handler {

	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	var mutex sync.Mutex
	var cpuValues []float64
	var netValues []uint64

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		testID := r.Header.Get("x-test")
		reqID := r.Header.Get("x-req")
		r = r.WithContext(context.WithValue(r.Context(), "x-test", testID))
		r = r.WithContext(context.WithValue(r.Context(), "x-req", reqID))

		if os.Getenv("ENV") != "production" {
			next.ServeHTTP(w, r)
			return
		}

		var cpuUsage float64
		var netUsage uint64

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

		mutex.Lock()
		cpuValues = append(cpuValues, cpuUsage)
		var cpuMed float64
		var sum float64
		for _, v := range cpuValues {
			sum += v
		}
		cpuMed = sum / float64(len(cpuValues))
		mutex.Unlock()

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
		mutex.Lock()
		netValues = append(netValues, netUsage)
		var netMed float64
		var sumNet uint64
		for _, v := range netValues {
			sumNet += v
		}
		netMed = float64(sumNet) / float64(len(netValues))
		mutex.Unlock()

		_, err = db.Exec(`
			INSERT INTO benchmark(test, resource, x, y)
			VALUES (?, ?, ?, ?), (?, ?, ?, ?)
		`,
			testID, "cpu", reqID, cpuMed,
			testID, "net", reqID, netMed/128.0,
		)
		if err != nil {
			log.Fatal(err)
		}

		next.ServeHTTP(w, r)
	})
}
