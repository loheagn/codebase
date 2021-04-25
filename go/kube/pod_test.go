package kube

import (
	"testing"
)

func Test_getPodLabel(t *testing.T) {
	type args struct {
		namespace string
		podName   string
	}
	tests := []struct {
		name string
		args args
	}{
		{args: args{"0-1", "firln-deployment-6cd4f4c776-nld84"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getPodLabel(tt.args.namespace, tt.args.podName)
		})
	}
}

func Test_listPod(t *testing.T) {
	type args struct {
		namespace string
		labels    map[string]string
	}
	tests := []struct {
		name string
		args args
	}{
		{args: args{"0-1", map[string]string{"commit": "9a50aa80144d91f6a1587c8987aeb69d6a3f7cd4"}}},
		{args: args{"0-1", map[string]string{"commit": "hahhahahahah"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listPod(tt.args.namespace, tt.args.labels)
		})
	}
}
