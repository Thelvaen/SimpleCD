package config

import (
	"math/rand"
	"time"

	"github.com/spf13/viper"
)

type Configuration struct {
	// HttpConf struct holds the HTTP server details
	HttpHost struct {
		Port        int
		Ssl         bool
		SslKeyPath  string
		SslCertPath string
		CSRF        []byte
		HashKey     []byte
		BlockKey    []byte
	}

	// MailConf struct holds the SMTP server details
	MailServer struct {
		Host     string
		Port     string
		From     string
		Username string
		Password string
	}

	// Repos is a slice of Repo
	Repos []struct {
		Title      string
		RepoType   string
		RemoteRepo string
		LocalRepo  string
		Key        string
		NotifyList []string
		Scenario   map[string]string
		Test       map[string]string
		TestRule   string
		Rollback   map[string]string
	}
}

var (
	Config Configuration

	// internal vars
	seededRand *rand.Rand
)

func init() {
	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	seededRand.Uint64()

	// Init Viper
	viper.SetConfigName("alambic")
	viper.AddConfigPath("/etc/alambic/")
	viper.AddConfigPath("/etc/")
	viper.AddConfigPath(".")

	// Loading conf in Viper
	err := viper.ReadInConfig()
	if err != nil {
		panic("can't load configuration")
	}
}
