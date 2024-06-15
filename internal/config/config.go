package config

import (
	"errors"
	"os"

	"github.com/st/internal/utils"
)

const (
	LDAPUSER  = "LDAPUSER"
	ECHO_FILE = "ECHO_FILE"
	ECHO_LIST = "ECHO_LIST"
)

type Config struct {
	User     string
	EchoFile string
	List     string
}

func NewConfig() (*Config, error) {
	user := os.Getenv(LDAPUSER)
	if user == "" {
		return nil, errors.New("LDAPUSER not exists")
	}

	f := os.Getenv(ECHO_FILE)
	if f == "" {
		return nil, errors.New("ECHO_FILE not exists")
	}
	if err := utils.CheckFile(f); err != nil {
		return nil, err
	}

	lst := os.Getenv(ECHO_LIST)
	if lst == "" {
		return nil, errors.New("ECHO_LIST not exists")
	}

	return &Config{User: user, EchoFile: f, List: lst}, nil
}
