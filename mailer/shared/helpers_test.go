package helpers

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	type args struct {
		key      string
		fallback string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Set Environemnt returns test variable",
			args: args{
				key:      "FOR_TEST",
				fallback: "Wouldntwork",
			},
			want: "wouldWork",
		},
		{
			name: "Test Set Environemnt returns test variable",
			args: args{
				key:      "NO_SUCH_VARIABLE",
				fallback: "TheFallbackWorks",
			},
			want: "TheFallbackWorks",
		},
	}
	os.Setenv("FOR_TEST", "wouldWork")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnv(tt.args.key, tt.args.fallback); got != tt.want {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
