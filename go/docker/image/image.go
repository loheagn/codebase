package image

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/archive"
	"github.com/loheagn/go-codebase/docker"
)

func Build(dockerFilepath, ctxPath string, tags []string, output io.Writer) error {
	cli, err := docker.GetClient()
	if err != nil {
		return err
	}
	buildOpts := types.ImageBuildOptions{
		Dockerfile: dockerFilepath,
		Tags:       tags,
	}
	buildCtx, _ := archive.TarWithOptions(ctxPath, &archive.TarOptions{})
	ctx := context.Background()

	resp, err := cli.ImageBuild(ctx, buildCtx, buildOpts)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(output, resp.Body)
	return err
}
