package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	//"github.com/otoolep/hraftd/http"
	//"github.com/otoolep/hraftd/store"

	//"github.com/hashicorp/raft"
	//"github.com/coreos/etcd/embed"
	//"github.com/coreos/etcd/client"
	"github.com/satori/go.uuid"

	"github.com/hartfordfive/distribd/httpd"
	"github.com/hartfordfive/distribd/worker"
	"github.com/hartfordfive/distribd/worker/web"
)

// Command line defaults
const (
	DefaultDataDir = "_distribd-data"
)

// Command line parameters
var (
	dataDir, joinAddr string
)

func init() {
	flag.StringVar(&dataDir, "data", DefaultDataDir, "Set Raft data directory")
	flag.StringVar(&joinAddr, "join", "", "Set join address, if any")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	fmt.Printf("Data Directory: %s\n", dataDir)
	if joinAddr != "" {
		fmt.Printf("Join Cluster at: %s\n", joinAddr)
	}
	fmt.Println("")

	if dataDir == "" {
		fmt.Fprintf(os.Stderr, "No Raft storage directory specified\n")
		os.Exit(1)
	}

	wn := web.NewWebNode()

}
