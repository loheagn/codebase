package docker

import (
	"bufio"
	"io"

	"github.com/docker/docker/client"
)

func GetClient() (cli *client.Client, err error) {
	return client.NewClientWithOpts(client.FromEnv)
}

func ReadOutput(rd io.Reader, output io.Writer) (err error) {
	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		_, _ = output.Write(scanner.Bytes())
		_, _ = output.Write([]byte{'\n'})
	}
	return
}
