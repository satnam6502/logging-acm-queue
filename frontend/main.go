// A synthetic front-end server program that acts as a proxy
// to a redis cluster and emits random log messages. Written to
// support an article on cluster level logging with Kubernetes.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/xyproto/simpleredis"
)

var (
	redisMaster = flag.String("redis-master", "redis-master:6379", "address of redis master")
	redisSlave  = flag.String("redis-slave", "redis-slave:6379", "address of redis slave")
	port        = flag.String("port", ":80", "port to listen on")
)

var (
	masterPool *simpleredis.ConnectionPool
	slavePool  *simpleredis.ConnectionPool
)

// ListRangeHandler lists the values for a key.
func ListRangeHandler(rw http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["key"]
	list := simpleredis.NewList(slavePool, key)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Printf("Slow read for key %v: %d ms\n", key, r1.Intn(100)+100)
	members := HandleError(list.GetAll()).([]string)
	membersJSON := HandleError(json.MarshalIndent(members, "", "  ")).([]byte)
	rw.Write(membersJSON)
}

// ListPushHandler adds a key/value.
func ListPushHandler(rw http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["key"]
	value := mux.Vars(req)["value"]
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Printf("Slow write for key %v: %d ms\n", key, r1.Intn(100)+100)
	list := simpleredis.NewList(masterPool, key)
	HandleError(nil, list.Add(value))
	ListRangeHandler(rw, req)
}

// InfoHandler reports the status of the redis store.
func InfoHandler(rw http.ResponseWriter, req *http.Request) {
	info := HandleError(masterPool.Get(0).Do("INFO")).([]byte)
	rw.Write(info)
}

// HandleError throws a panic if somethign goes wrong.
func HandleError(result interface{}, err error) (r interface{}) {
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	flag.Parse()
	fmt.Printf("redis master: %s\n", *redisMaster)
	fmt.Printf("redis slave: %s\n", *redisSlave)
	masterPool = simpleredis.NewConnectionPoolHost(*redisMaster)
	defer masterPool.Close()
	slavePool = simpleredis.NewConnectionPoolHost(*redisSlave)
	defer slavePool.Close()

	r := mux.NewRouter()
	r.Path("/lrange/{key}").Methods("GET").HandlerFunc(ListRangeHandler)
	r.Path("/rpush/{key}/{value}").Methods("GET").HandlerFunc(ListPushHandler)
	r.Path("/info").Methods("GET").HandlerFunc(InfoHandler)

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(*port)
}
