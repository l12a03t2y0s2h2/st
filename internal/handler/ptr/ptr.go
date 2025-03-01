package ptr

import (
	"fmt"
	"net"
	"os"
	"slices"
	"strings"
)

const ARPA = ".in-addr.arpa"
const ARGS = 3 // |0| binary |1| ptr |2| 123.45.67.89

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle() (string, error) {
	if len(os.Args) != ARGS {
		return "", fmt.Errorf("args number error [expected: %d]", ARGS)
	}

	// Check IP
	var ipV4 net.IP
	if ipV4 = net.ParseIP(os.Args[2]); ipV4 == nil {
		return "", fmt.Errorf("not correct IP [provided: %s]", ipV4.String())
	}

	// IP job
	octets := strings.Split(ipV4.String(), ".")
	slices.Reverse(octets)

	return fmt.Sprintf("%s%s", strings.Join(octets, "."), ARPA), nil
}
