package internal

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func CreateConfigDir() {

	if _, err := os.Stat(".glc"); os.IsNotExist(err) {
		err := os.Mkdir(".glc", os.ModePerm)
		if err != nil {
			log.Errorln("Unable to make the config directory %s", err)
		}
	}
	
	_, err := os.Create(".glc/config.yaml")
	if err != nil {
		log.Errorln("Unable to create the config file %s", err)
	}
}
