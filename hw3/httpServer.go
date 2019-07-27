package main

import (
	"fmt"
	"github.com/jhkolb/cs162-hw3/http"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
)

const usageMsg = "Usage: ./httpserver --files www_directory/ --port 8000 [--num-threads 5]\n" +
	"       ./httpserver --proxy inst.eecs.berkeley.edu:80 --port 8000 [--num-threads 5]\n"
const maxQueueSize = 50

var (
	serverFilesDirectory string
	proxyAddress         string
	proxyPort            int

	interrupted chan os.Signal
)

func handleFilesRequest(connection net.Conn) {
    // TODO Fill this in to complete Task #2
}

func handleProxyRequest(clientConn net.Conn) {
    // TODO Fill this in to complete Task #3
    // Open a connection to the specified upstream server
    // Create two goroutines to forward traffic
    // Use sync.WaitGroup to block until the goroutines have finished
}

func handleSigInt() {
}

func initWorkerPool(numThreads int, requestHandler func(net.Conn)) {
    // TODO Fill this in as part of Task #1
    // Create a fixed number of goroutines to handle requests
}

func serveForever(numThreads int, port int, requestHandler func(net.Conn)) {
    // TODO Fill this in as part of Task #1
    // Create a Listener and accept client connections
    // Pass connections to the thread pool via a channel
}

func exitWithUsage() {
	fmt.Fprintf(os.Stderr, usageMsg)
	os.Exit(-1)
}

func main() {
	// Command line argument parsing
	var requestHandler func(net.Conn)
	var serverPort int
	numThreads := 1
	var err error

	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--files":
			requestHandler = handleFilesRequest
			if i == len(os.Args)-1 {
				fmt.Fprintln(os.Stderr, "Expected argument after --files")
				exitWithUsage()
			}
			serverFilesDirectory = os.Args[i+1]
			i++

		case "--proxy":
			requestHandler = handleProxyRequest
			if i == len(os.Args)-1 {
				fmt.Fprintln(os.Stderr, "Expected argument after --proxy")
				exitWithUsage()
			}
			proxyTarget := os.Args[i+1]
			i++

			tokens := strings.SplitN(proxyTarget, ":", 2)
			proxyAddress = tokens[0]
			proxyPort, err = strconv.Atoi(tokens[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Expected integer for proxy port")
				exitWithUsage()
			}

		case "--port":
			if i == len(os.Args)-1 {
				fmt.Fprintln(os.Stderr, "Expected argument after --port")
				exitWithUsage()
			}

			portStr := os.Args[i+1]
			i++
			serverPort, err = strconv.Atoi(portStr)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Expected integer value for --port argument")
				exitWithUsage()
			}

		case "--num-threads":
			if i == len(os.Args)-1 {
				fmt.Fprintln(os.Stderr, "Expected argument after --num-threads")
				exitWithUsage()
			}
			numThreadsStr := os.Args[i+1]
			i++
			numThreads, err = strconv.Atoi(numThreadsStr)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Expected positive integer value for --num-threads argument")
				exitWithUsage()
			}

		case "--help":
			fmt.Printf(usageMsg)
			os.Exit(0)

		default:
			fmt.Fprintf(os.Stderr, "Unexpected command line argument %s\n", os.Args[i])
			exitWithUsage()
		}
	}

	if requestHandler == nil {
		fmt.Fprintln(os.Stderr, "Must specify one of either --files or --proxy")
		exitWithUsage()
	}

	// Set up a handler for SIGINT, used in Task #4
	interrupted = make(chan os.Signal, 1)
	signal.Notify(interrupted, os.Interrupt)
	serveForever(numThreads, serverPort, requestHandler)
}

