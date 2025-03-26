package main

import (
	"context"
	"io"
	"os"
	"fmt"
	"github.com/docker/docker/api/types/container"
        "github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func main() {
    var name string

    fmt.Println("Enter a name for the container you wish to create :")
    fmt.Scanf("%s\n", &name)

    ctx := context.Background()
    cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
    if err != nil {
        panic(err)
    }
    defer cli.Close()

    reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", image.PullOptions{})
    if err != nil {
        panic(err)
    }
    io.Copy(os.Stdout, reader)

    resp, err := cli.ContainerCreate(ctx, &container.Config{
        Image: "alpine",
        Cmd:   []string{"echo", "hello world"},
    }, nil, nil, nil, name)
    if err != nil {
        panic(err)
    }

    if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
        panic(err)
    }

    statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
    select {
    case err := <-errCh:
        if err != nil {
            panic(err)
        }
    case <-statusCh:
    }

    out, err := cli.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true})
    if err != nil {
        panic(err)
    }

    stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}
