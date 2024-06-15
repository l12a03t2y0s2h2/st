package ea

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/st/internal/config"
)

const SCP = "scp"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle() (string, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return "", err
	}

	var wg sync.WaitGroup
	start := time.Now()
	for _, e := range strings.Split(cfg.List, ",") {
		e = strings.TrimSpace(strings.ReplaceAll(e, "echo", "e"))
		wg.Add(1)
		go func(e string) {
			defer wg.Done()
			cmd := exec.Command(SCP, cfg.EchoFile, fmt.Sprintf("%s:/home/%s", e, cfg.User))
			cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
			if err := cmd.Run(); err != nil {
				fmt.Printf("[%s] not ok\n\n", e)
				return
			}
			fmt.Printf("[%s] ok\n\n", e)
		}(e)
	}
	wg.Wait()

	return fmt.Sprintf("Done. Elapsed time: %.2f sec.", time.Since(start).Seconds()), nil
}
