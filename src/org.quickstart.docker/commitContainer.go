package main

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {

	//Commit a container
	//Commit a container to create an image from its contents:

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

	createResp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"touch", "/helloworld"},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, createResp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, createResp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	commitResp, err := cli.ContainerCommit(ctx, createResp.ID, types.ContainerCommitOptions{Reference: "helloworld"})
	if err != nil {
		panic(err)
	}

	fmt.Println(commitResp.ID)
}
