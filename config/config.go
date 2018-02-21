package config

import (
	"os"

	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/google/uuid"
)

//default values for application
const (
	//api const
	DefaultAPIPort       = "8000"
	DefaultAPIIP         = "0.0.0.0"
	DefaultAPIServiceURL = "http://api-svc"
	//basic auth enabled or not
	DefaultBasicAuthentication = false
	DefaultTimeBomb            = false
	DefaultAPIPassword         = "default"

	//frontend defaults
	DefaultFEPort = "8080"
	DefaultFEIP   = "0.0.0.0"

	//ENV
	EnvAPIPort       = "API_PORT"
	EnvAPIIP         = "API_IP"
	EnvFEPort        = "FE_PORT"
	EnvFEIP          = "FE_IP"
	EnvBasicAuth     = "API_BASIC_AUTH"
	EnvAPIServiceURL = "API_SVC"
	EnvAPIPassword   = "API_PASSWORD"

	EnvTimeBomb = "TIME_BOMB"
)

//Options structures for application default and configuration
type Options struct {
	Default     string
	Environment string
}

// GenerateID for Note
func GenerateID() (id string) {
	return uuid.New().String()
}

// Get - gets specified variable from either environment or default one
func Get(variable string) string {

	var config = map[string]Options{
		"EnvAPIPort": {
			Default:     DefaultAPIPort,
			Environment: EnvAPIPort,
		},
		"EnvAPIIP": {
			Default:     DefaultAPIIP,
			Environment: EnvAPIIP,
		},
		"EnvBasicAuth": {
			Default:     strconv.FormatBool(DefaultBasicAuthentication),
			Environment: EnvBasicAuth,
		},
		"EnvFEPort": {
			Default:     DefaultFEPort,
			Environment: EnvFEPort,
		},
		"EnvFEIP": {
			Default:     DefaultFEIP,
			Environment: EnvFEIP,
		},
		"EnvAPIServiceURL": {
			Default:     DefaultAPIServiceURL,
			Environment: EnvAPIServiceURL,
		},
		"EnvAPIPassword": {
			Default:     DefaultAPIPassword,
			Environment: EnvAPIPassword,
		},
		"EnvTimeBomb": {
			Default:     strconv.FormatBool(DefaultTimeBomb),
			Environment: EnvTimeBomb,
		},
	}

	for k, v := range config {
		if k == variable {
			if os.Getenv(v.Environment) != "" {
				log.WithFields(log.Fields{
					"key":   k,
					"value": v.Environment,
				}).Debug("config: setting configuration")
				return os.Getenv(v.Environment)
			}
			log.WithFields(log.Fields{
				"key":   k,
				"value": v.Default,
			}).Debug("config: setting configuration")
			return v.Default

		}
	}
	return ""
}
