package target

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"os/exec"
)

const (
	ARGS  = 4 // |0| binary |1| 443/80 |2| domain |3| target_ip
	P443  = "443"
	HTTPS = "https://"
	HTTP  = "http://"
	CURL  = "curl"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle() (string, error) {
	if len(os.Args) != ARGS {
		return "", fmt.Errorf("args number error [expected: %d]", ARGS)
	}

	// Get port
	port := os.Args[1]
	// Get protocol
	protocol := map[bool]string{true: HTTPS, false: HTTP}[port == P443]
	// Get domain
	domain, err := url.Parse(protocol + os.Args[2])
	if err != nil {
		return "", fmt.Errorf("not correct domain [provided: %s]", os.Args[2])
	}

	// Get IP
	ip := net.ParseIP(os.Args[3])
	if ip == nil {
		return "", fmt.Errorf("not correct IP [provided: %s]", os.Args[3])
	}

	// Collect args for curl
	args := []string{
		"--resolve", // --resolve <[+]host:port:addr[,addr]...> Provide a custom address for a specific host and port pair
		fmt.Sprintf("%s:%s:%s", domain.Host, port, ip),
		"-o", // -o, --output <file> Write to file instead of stdout
		"/dev/null",
		"-k",        // -k, --insecure (TLS  SFTP  SCP) This option makes curl skip the verification step and proceed without checking
		"-v",        // -v, --verbose Make the operation more talkative
		"--http1.1", // --http1.1 (HTTP) Tells curl to use HTTP version 1.1
		"-w",
		"'Connect: %{time_connect} TTFB: %{time_starttransfer} Total time: %{time_total}'", // time_total The total time, in seconds, that the full operation lasted
		fmt.Sprintf("%s://%s%s", domain.Scheme, domain.Host, domain.Path),
	}

	cmd := exec.Command(CURL, args...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err = cmd.Run(); err != nil {
		return "", fmt.Errorf("unable to run curl command")
	}

	return fmt.Sprintf("\n%s\n", cmd.String()), nil
}
