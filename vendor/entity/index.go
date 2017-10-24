package entity

import (
	"os"
	"utils"

	"github.com/spf13/viper"
)

type model interface {
	Init(string)
	load()
	dump()
}

type modelConfig struct {
	model    model
	filename string
}

var (
	logger = utils.NewLogger()
	models []modelConfig
)

func addModel(model model, filename string) {
	models = append(models, modelConfig{
		model:    model,
		filename: filename,
	})
}

// Init initializes registered models
func Init() {
	finished := make(chan bool)
	// initialize all models concurrently
	for _, m := range models {
		go func(m modelConfig) {
			path := viper.GetString(m.filename)
			if len(path) == 0 {
				os.Mkdir("data", 0755)
				path = "data/" + m.filename
			}
			m.model.Init(path)
			finished <- true
		}(m)
	}
	// wait for all models to finish initialization
	for _ = range models {
		<-finished
	}
}
