package config_test

import (
	"os"
	"testing"
	"time"

	"github.com/SergioAn2003/go-config"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	r := require.New(t)

	testCfg, err := config.New("example.env")
	r.NoError(err)

	tests := []struct {
		name       string
		envPath    string
		wantConfig config.Config
		wantErr    error
	}{
		{
			name:       "succes",
			envPath:    "example.env",
			wantConfig: testCfg,
			wantErr:    nil,
		},
		{
			name:       "fail path",
			envPath:    "sfjskd",
			wantConfig: nil,
			wantErr:    config.ErrInvalidConfigPath,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := config.New(tt.envPath)

			if tt.wantErr != nil {
				r.ErrorIs(err, tt.wantErr)
			} else {
				r.Nil(err)
			}

			if tt.wantConfig != nil {
				r.Implements((*config.Config)(nil), cfg)
			} else {
				r.Nil(cfg)
			}
		})
	}
}

func TestGetters(t *testing.T) {
	r := require.New(t)

	cfg, err := config.New("example.env")
	r.NoError(err)

	tests := []struct {
		name   string
		setEnv func()
		getFn  func() any
		want   func() any
	}{
		{
			name:   "string",
			setEnv: func() { r.NoError(os.Setenv("TEST_STRING", "test")) },
			getFn:  func() any { return cfg.String("TEST_STRING") },
			want:   func() any { return "test" },
		},
		{
			name:   "string slice",
			setEnv: func() { r.NoError(os.Setenv("TEST_STRING_SLICE", "test test1")) },
			getFn:  func() any { return cfg.StringSlice("TEST_STRING_SLICE") },
			want:   func() any { return []string{"test", "test1"} },
		},
		{
			name:   "int",
			setEnv: func() { r.NoError(os.Setenv("TEST_INT", "1")) },
			getFn:  func() any { return cfg.Int("TEST_INT") },
			want:   func() any { return 1 },
		},
		{
			name:   "float",
			setEnv: func() { r.NoError(os.Setenv("TEST_FLOAT", "1.1")) },
			getFn:  func() any { return cfg.Float("TEST_FLOAT") },
			want:   func() any { return 1.1 },
		},
		{
			name:   "bool",
			setEnv: func() { r.NoError(os.Setenv("TEST_BOOL", "true")) },
			getFn:  func() any { return cfg.Bool("TEST_BOOL") },
			want:   func() any { return true },
		},
		{
			name:   "time",
			setEnv: func() { r.NoError(os.Setenv("TEST_TIME", "2022-01-01T00:00:00Z")) },
			getFn:  func() any { return cfg.Time("TEST_TIME") },
			want:   func() any { return time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC) },
		},
		{
			name:   "duration",
			setEnv: func() { r.NoError(os.Setenv("TEST_DURATION", "1h1m1s")) },
			getFn:  func() any { return cfg.Duration("TEST_DURATION") },
			want:   func() any { return time.Hour + time.Minute + time.Second },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setEnv()
			r.Equal(tt.want(), tt.getFn())
		})
	}
}
