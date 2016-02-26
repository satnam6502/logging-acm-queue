// A synthetic front-end server program that acts as a proxy
// to a redis cluster and emits random log messages. Written to
// support an article on cluster level logging with Kubernetes.

package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	keyrange = 1000
)

var (
	redisMaster = flag.String("redis-master", "redis-master:6379", "address of redis master")
	redisSlave  = flag.String("redis-slave", "redis-slave:6379", "address of redis slave")
	frontend    = flag.String("frontend", ":3000", "frontend server")
	requests    = flag.Int("requests", 10000, "number of requests to perform")
)

func main() {
	flag.Parse()
	client := &http.Client{}
	// Make some writes
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for *requests > 0 {
		if r1.Intn(10) < 2 {
			// perform a write
			key := r1.Intn(keyrange)
			value := r1.Intn(1000000)
			resp, err := client.Get(fmt.Sprintf("http://%s/rpush/k%d/v%d", *frontend, key, value))
			if err != nil {
				log.Printf("Write failed: %v", err)
			}
			resp.Body.Close()
		} else {
			// perform a read
			key := r1.Intn(keyrange)
			resp, err := client.Get(fmt.Sprintf("http://%s/lrange/k%d", *frontend, key))
			if err != nil {
				log.Printf("Read failed: %v", err)
			}
			resp.Body.Close()
		}
		// Take a pause
		time.Sleep(time.Duration(50+r1.Intn(1000)) * time.Millisecond)
		*requests--
	}
}
