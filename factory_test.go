package xinconf_test

import (
	"testing"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/gz4z2b/xinconf"
	"github.com/gz4z2b/xinconf/types"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func TestViper(t *testing.T) {
	type Conf struct {
		Host string `yaml:"host"`
	}
	viper := viper.New()
	viper.AddConfigPath("conf")
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	require.NoError(t, err)
	var conf Conf
	err = viper.UnmarshalKey("db.mysql", &conf)
	t.Log(conf, err)
}

func TestNewXinConf(t *testing.T) {
	type args struct {
		conf types.XinConfConf
		key  string
	}
	type Result struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr error
	}{
		{
			name: "正常",
			args: args{
				conf: types.XinConfConf{
					ConfName:     "dev",
					ConfType:     types.ConfTypeViper,
					ConfFileType: "yaml",
					ConfPath:     []string{"conf"},
				},
				key: "db.mysql",
			},
			want: Result{
				Host: "127.0.0.1",
				Port: "3306",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := xinconf.NewXinConf(tt.args.conf)
			require.NoError(t, err)
			var result Result
			err = got.Unmarshal(tt.args.key, &result)
			assert.Equal(t, tt.want, result)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestChangeConf(t *testing.T) {
	type Conf struct {
		Host string `yaml:"host"`
	}
	viper := viper.New()
	viper.AddConfigPath("conf")
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	require.NoError(t, err)
	viper.OnConfigChange(func(in fsnotify.Event) {
		t.Log("配置文件发生变更")
		t.Log(in)
	})
	server := gin.Default()
	server.Run(":18081")
}
