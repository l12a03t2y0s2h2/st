package hosts

import (
	"bufio"
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
)

const ARGS = 4 // |0| binary |1| hosts |2| domain |3| target_ip
const HOSTS = "/etc/hosts"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle() (string, error) {
	if len(os.Args) != ARGS {
		return "", fmt.Errorf("args number error [expected: %d]", ARGS)
	}

	// Get domain
	domain, err := url.Parse(os.Args[2])
	if err != nil {
		return "", fmt.Errorf("not correct domain [provided: %s]", os.Args[2])
	}

	// Get IP
	ip := net.ParseIP(os.Args[3])
	if ip == nil {
		return "", fmt.Errorf("not correct IP [provided: %s]", os.Args[3])
	}

	// Create record
	record := fmt.Sprintf("%s %s", ip, domain)

	// https://pkg.go.dev/os#O_RDONLY
	f, err := os.OpenFile(HOSTS, os.O_RDWR, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// File job
	records := make([]string, 0)
	var exists bool
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		recordFromFile := scanner.Text()
		if recordFromFile != record {
			records = append(records, recordFromFile)
		} else {
			exists = true
		}
	}
	if err = scanner.Err(); err != nil {
		return "", err
	}

	result := "Deleted"
	if !exists {
		records = append(records, record)
		result = "Added"
	}

	// Clean file and set cursor for 0 position
	if err = f.Truncate(0); err != nil {
		return "", err
	}
	if _, err = f.Seek(0, 0); err != nil {
		return "", err
	}

	// Write records to file
	if _, err = f.WriteString(strings.Join(records, "\n") + "\n"); err != nil {
		return "", err
	}

	return result, nil
}
