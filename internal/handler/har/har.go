package har

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
)

const ARGS = 5 // |0| binary |1| har |2| path_to_har |3| target_ip |4| echo_ip
const BASE_STRING = "curl --resolve %s:443:%s --interface %s --http1.1 -k -v -o /dev/null"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle() (string, error) {
	if len(os.Args) != ARGS {
		return "", fmt.Errorf("args number error [expected: %d]", ARGS)
	}

	targetIp := os.Args[3]
	if ipInst := net.ParseIP(targetIp); ipInst == nil {
		return "", fmt.Errorf("not correct IP [provided: %s]", targetIp)
	}

	echoIp := os.Args[4]
	if ipInst := net.ParseIP(targetIp); ipInst == nil {
		return "", fmt.Errorf("not correct IP [provided: %s]", echoIp)
	}

	// Read file
	data, err := os.ReadFile(os.Args[2])
	if err != nil {
		return "", err
	}
	dataStr := strings.TrimSpace(string(data))

	// Check browser engine
	if !strings.Contains(dataStr, "Chromium") {
		return "", fmt.Errorf("browser error")
	}

	// Get records from data
	records := strings.Split(dataStr, "\n")
	for i, record := range records {
		records[i] = strings.Trim(record, "\\ ")
	}

	u, err := url.Parse(strings.Trim(strings.Split(records[0], " ")[1], "'"))
	if err != nil {
		return "", fmt.Errorf("url.Parse error > %s", err)
	}

	bestString := fmt.Sprintf(BASE_STRING, u.Host, targetIp, echoIp)

	return fmt.Sprintf("%s %s %s", bestString, strings.Join(records[1:], " "), u.String()), nil
}
