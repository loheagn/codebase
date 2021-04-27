package xorm_demo

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

func Test_initEngine(t *testing.T) {
	tests := []struct {
		name    string
		want    *xorm.Engine
		wantErr bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := initEngine()
			if (err != nil) != tt.wantErr {
				t.Errorf("initEngine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("initEngine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderDemo(t *testing.T) {
	t.Run("desc", func(t *testing.T) {
		if got := OrderDemo("desc"); got.Number != 22 {
			t.Errorf("OrderDemo() = %v, want %v", got, 22)
		}
	})
	t.Run("asc", func(t *testing.T) {
		if got := OrderDemo("asc"); got.Number != 20 {
			t.Errorf("OrderDemo() = %v, want %v", got, 20)
		}
	})
}
