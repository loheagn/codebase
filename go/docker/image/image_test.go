package image

import (
	"bytes"
	"testing"
)

func Test_buildImage(t *testing.T) {
	type args struct {
		dockerFilepath string
		ctxPath        string
		tags           []string
	}
	tests := []struct {
		name          string
		args          args
		wantOutput    string
		wantOutputErr string
		wantErr       bool
	}{
		{args: args{"Dockerfile", "../example/ubuntu-test", []string{"test/ubuntu:20.04"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := &bytes.Buffer{}
			err := Build(tt.args.dockerFilepath, tt.args.ctxPath, tt.args.tags, output)
			t.Error(output.String())
			if (err != nil) != tt.wantErr {
				t.Errorf("buildImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
