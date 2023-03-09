package Config

import (
	"testing"
)

func TestApp_Inc(t *testing.T) {
	a := &App{
		Mp:   make(map[string]int),
		Size: 0,
	}
	type args struct {
		id string
	}
	tests := []struct {
		name     string
		args     args
		wantdata int
	}{
		{
			name: "Test1",
			args: args{
				id: "a",
			},
			wantdata: 1,
		},
		{
			name: "Test2",
			args: args{
				id: "a",
			},
			wantdata: 2,
		},
		{
			name: "Test3",
			args: args{
				id: "b",
			},
			wantdata: 1,
		},
	}
	for _, tt := range tests {
		a.Inc(tt.args.id)
		if a.V[a.Mp[tt.args.id]].Views != tt.wantdata {
			t.Error("Error in incrementing views")
		}
	}
}

func TestApp_Get(t *testing.T) {
	a := &App{
		Mp:   make(map[string]int),
		Size: 0,
	}
	a.Inc("a")
	a.Inc("a")
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{
				id: "a",
			},
			want: 2,
		},
		{
			name: "Test2",
			args: args{
				id: "b",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		if got := a.Get(tt.args.id); got != tt.want {
			t.Errorf("App.Get() = %v, want %v", got, tt.want)
		}

	}
}