package main

import (
	"os"
	"strconv"

	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/bobbydeveaux/go-example-app/config"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {

	bomb, _ := strconv.ParseBool(config.Get("EnvTimeBomb"))
	if bomb {
		go timeBomb()

	}

	log.Info("Starting OCP Demo")
	NewRouter()
}

func timeBomb() {
	log.Error("Time Bomb in 5 Seconds")
	time.Sleep(5000 * time.Millisecond)
	log.Fatal("Time Bomb!!!")
}
