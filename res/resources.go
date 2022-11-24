// Package res provide global resources in order to separate the essential
// config and general values from the `bot.yaml`
package res

import (
	"encoding/json"
	"io"
	"os"

	"github.com/sslime336/awbot/config"
	"github.com/sslime336/awbot/utils"
)

// R represents the global resources' instance.
var R *Entity

// M is a map represents the R for general query useage.
var M *EntityMap

func Load() {
	f, err := os.Open(config.ResourcesFilePath)
	utils.Check(err)

	data, err := io.ReadAll(f)
	utils.Check(err)

	// build resources entity
	R = new(Entity)
	err = json.Unmarshal(data, R)
	utils.Check(err)

	mapping()
}

func mapping() {
	if R == nil {
		return
	}

	lazyInitM()

	mapString()
	mapNumber()
	mapGroup()
	mapCron()
}

func lazyInitM() {
	if M != nil {
		return
	}
	M = new(EntityMap)
	M.String = make(map[string]string, 5)
	M.Number = make(map[string]int64, 5)
	M.Group = make(map[string]int64, 3)
	M.Cron = make(map[string]string, 5)
}

func mapCron() {
	for _, cron := range R.Crons {
		M.Cron[cron.ID] = cron.Cron
	}
}

func mapGroup() {
	for _, group := range R.Groups {
		M.Group[group.ID] = group.Code
	}
}

func mapNumber() {
	for _, num := range R.Values.Numbers {
		M.Number[num.ID] = num.Value
	}
}

func mapString() {
	for _, str := range R.Values.Strings {
		M.String[str.ID] = str.Value
	}
}

type EntityMap struct {
	String map[string]string // [id]:value
	Number map[string]int64  // [id]:value
	Group  map[string]int64  // [id/name]:code
	Cron   map[string]string // [id]:cronStmt
}

type Entity struct {
	Values Values  `json:"values"`
	Groups []Group `json:"groups"`
	Crons  []Cron  `json:"crons"`
}

type Cron struct {
	ID   string `json:"id"`
	Cron string `json:"cron"`
}

type Group struct {
	ID   string `json:"id"`
	Code int64  `json:"code"`
}

type Values struct {
	Strings []String `json:"strings"`
	Numbers []Number `json:"numbers"`
}

type Number struct {
	ID    string `json:"id"`
	Value int64  `json:"value"`
}

type String struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}
