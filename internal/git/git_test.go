package git

import (
	"testing"
)

func TestGit(t *testing.T) {

}

func TestRebase(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rebase()
		})
	}
}

func TestCurBranch(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "bad test",
			want:    "jgr142/feat/gitTools",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CurBranch()
			if (err != nil) != tt.wantErr {
				t.Errorf("CurBranch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CurBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}
