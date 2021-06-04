package container

import (
	"bytes"
	"testing"

	"github.com/loheagn/go-codebase/docker/image"
)

func Test_Run(t *testing.T) {
	const tag string = "test/ubuntu:20.04"
	// 先创建一下这个测试用的基础镜像
	output := &bytes.Buffer{}
	_ = image.Build("./Dockerfile", "../example/ubuntu-test", []string{tag}, output)

	type args struct {
		image  string
		config *RunOption
	}
	tests := []struct {
		name         string
		args         args
		wantExitCode int
		wantErr      bool
	}{
		{args: args{image: tag, config: &RunOption{Image: tag, Cmd: []string{"./error.sh"}}}, wantExitCode: 1},
		{args: args{image: tag, config: &RunOption{Image: tag, Cmd: []string{"./success.sh"}}}, wantExitCode: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if exitCode, err := Run(tt.args.config, output); (err != nil) != tt.wantErr || exitCode != tt.wantExitCode {
				t.Log(output)
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				t.Errorf("Run() exitCode = %v, wantExitCode %v", exitCode, tt.wantExitCode)
				return
			}
		})
	}
}
