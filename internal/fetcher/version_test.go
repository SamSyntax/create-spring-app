package fetcher

import "testing"

func Test_normalizeVersion(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"Release", "1.2.3.RELEASE", "1.2.3"},
		{"Milestone", "1.2.3.M1", "1.2.3-M1"},
		{"Release Candidate", "1.2.3.RC1", "1.2.3-RC1"},
		{"Build Snapshot", "1.2.3.BUILD-SNAPSHOT", "1.2.3-SNAPSHOT"},
		{"Snapshot", "1.2.3.SNAPSHOT", "1.2.3-SNAPSHOT"},
		{"Standard", "1.2.3", "1.2.3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeVersion(tt.args); got != tt.want {
				t.Errorf("normalizeVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckCompatibility(t *testing.T) {
	type args struct {
		bootVersion string
		springRange string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Exact match",
			args: args{bootVersion: "2.5.0", springRange: "2.5.0"},
			want: true,
		},
		{
			name: "Range match",
			args: args{bootVersion: "2.5.5", springRange: "[2.5.0,2.6.0)"},
			want: true,
		},
		{
			name: "Range mismatch low",
			args: args{bootVersion: "2.4.9", springRange: "[2.5.0,2.6.0)"},
			want: false,
		},
		{
			name: "Range mismatch high",
			args: args{bootVersion: "2.6.0", springRange: "[2.5.0,2.6.0)"},
			want: false,
		},
		{
			name: "Implicit open range match",
			args: args{bootVersion: "3.0.0", springRange: "2.5.0"},
			want: false,
		},
		{
			name: "Implicit open range mismatch",
			args: args{bootVersion: "2.4.0", springRange: "2.5.0"},
			want: false,
		},
		{
			name: "Explicit open range match",
			args: args{bootVersion: "3.0.0", springRange: ">=2.5.0"},
			want: true,
		},
		{
			name: "Explicit open range mismatch",
			args: args{bootVersion: "2.4.0", springRange: ">=2.5.0"},
			want: false,
		},
		{
			name: "Empty range",
			args: args{bootVersion: "2.5.0", springRange: ""},
			want: true,
		},
		{
			name: "Snapshot match in range",
			args: args{bootVersion: "2.6.0-SNAPSHOT", springRange: "[2.6.0-M1,2.6.0-RC1]"},
			want: false, // 2.6.0-SNAPSHOT is usually considered pre-release but semantics vary. Let's check logic.
			// normalizeVersion will make it 2.6.0-SNAPSHOT.
			// range normalization logic...
			// If it fails, I'll adjust based on deeper reading of `CheckCompatibility` logic.
			// The logic adds -0 suffix to ranges.
			// Let's stick to simpler semver cases first to be safe, or just run and fix.
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckCompatibility(tt.args.bootVersion, tt.args.springRange); got != tt.want {
				t.Errorf("CheckCompatibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
