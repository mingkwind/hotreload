package hotreload

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/json-iterator/go"
)

type config struct {
	Model string `json:"model"`
}

var cnf config

func loadConfig(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln("[TEST_SIGUSR] Load config: ", err)
	}

	err = jsoniter.Unmarshal(file, &cnf)
	if err != nil {
		log.Fatalln("[TEST_SIGUSR] Para config failed: ", err)
	}

	fmt.Println(cnf)

	return nil
}

func TestWatcher(t *testing.T) {
	// 注意：监控目录为conf
	// step 1
	Register("conf/config.json", loadConfig)
	Register("conf/app/config.json", loadConfig)

	// step 2
	Watcher()
}
