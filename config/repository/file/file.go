package file

import (
	"github.com/spf13/viper"

	"github.com/klinsmansun/zlassignment/model"
)

type fileConfigRepository struct {
	filePath string
	fileName string
	fileType string
}

func CreateConfigLoader(filePath, fileName, fileType string) model.ConfigLoader {
	loader := &fileConfigRepository{
		filePath: filePath,
		fileName: fileName,
		fileType: fileType,
	}

	return loader
}

func (f *fileConfigRepository) LoadConfig() (config *model.Config) {
	viper.AddConfigPath(f.filePath)
	viper.SetConfigType(f.fileType)
	viper.SetConfigName(f.fileName)
	viper.ReadInConfig()

	config = new(model.Config)
	viper.Unmarshal(config)
	return
}
