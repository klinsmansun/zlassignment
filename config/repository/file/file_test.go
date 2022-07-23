package file

import (
	"reflect"
	"testing"

	"github.com/klinsmansun/zlassignment/model"
)

func Test_fileConfigRepository_LoadConfig(t *testing.T) {
	type fields struct {
		fileName string
		fileType string
		filePath string
	}
	tests := []struct {
		name       string
		fields     fields
		orders     model.OrderRequest
		wantConfig *model.Config
	}{
		{
			name: "Test1",
			fields: fields{
				fileName: "config-test1",
				fileType: "yml",
				filePath: ".",
			},
			wantConfig: &model.Config{
				Version: "1.0",
				Log: model.LogConfig{
					Level: 7,
				},
				Core: model.CoreConfig{
					TradeInterval: 5,
					ChannelLength: 100,
				},
				GRPC: model.GRPCConfig{
					ListenIP:   "0.0.0.0",
					ListenPort: "80",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fileConfigRepository{
				fileName: tt.fields.fileName,
				fileType: tt.fields.fileType,
				filePath: tt.fields.filePath,
			}
			if gotConfig := f.LoadConfig(); !reflect.DeepEqual(gotConfig, tt.wantConfig) {
				t.Errorf("fileConfigRepository.LoadConfig() = %v, want %v", gotConfig, tt.wantConfig)
			}
		})
	}
}
