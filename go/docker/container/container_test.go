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
	outputErr := &bytes.Buffer{}
	_ = image.Build("./Dockerfile", "../example/ubuntu-test", []string{tag}, output)
	// outputS := output.String()
	type args struct {
		image  string
		config *RunOption
	}
	tests := []struct {
		name          string
		args          args
		wantOutputErr bool
		wantErr       bool
	}{
		{args: args{image: tag, config: &RunOption{Image: tag, Cmd: []string{"./error.sh"}}}},
		{args: args{image: tag, config: &RunOption{Image: tag, Cmd: []string{"./success.sh"}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Run(tt.args.config, output, outputErr); (err != nil) != tt.wantErr {
				t.Log(output)
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput := outputErr.String(); (outputErr.String() != "") != tt.wantOutputErr {
				t.Errorf("Run() = %v, want %v", gotOutput, tt.wantOutputErr)
			}
		})
	}
}
