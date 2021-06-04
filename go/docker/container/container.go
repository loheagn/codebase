package container

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/loheagn/go-codebase/docker"
)

type RunOption struct {
	Image string
	Cmd   []string
}

func Run(opt *RunOption, output io.Writer) (exitCode int, err error) {
	ctx := context.Background()
	cli, err := docker.GetClient()
	if err != nil {
		return 1, err
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: opt.Image,
		Cmd:   opt.Cmd,
	}, nil, nil, nil, "")
	if err != nil {
		return 1, err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return 1, err
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return 1, err
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return 1, err
	}

	_, err = stdcopy.StdCopy(output, output, out)
	if err != nil {
		return 1, err
	}

	status, err := cli.ContainerInspect(ctx, resp.ID)
	return status.State.ExitCode, err
}
