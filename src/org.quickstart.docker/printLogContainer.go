package main

import (
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {

	//Print the logs of a specific container

	/*ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}*/

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	options := types.ContainerLogsOptions{ShowStdout: true}
	// Replace this ID with a container that really exists
	out, err := cli.ContainerLogs(ctx, "f6105e4b6937", options)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)
}
