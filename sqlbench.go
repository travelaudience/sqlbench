/*
Package sqlbench can run a set of queries against a postgresql database and collect execution statistics.

Queries to run and how to run them is set in a json file that can must be passed.
*/
package sqlbench

import (
	"encoding/json"
	"io/ioutil"
)

// New will return a Bench structure that can be used to control the benchmark.
// configFile will point to a json file which contains the query and settings for benchmark.
// A sample config can be found in example directoy.
func New(configFile string) (*Bench, error) {
	b := Bench{}
	var err error
	if b.config, err = config(configFile); err != nil {
		return nil, err
	}

	b.runner = &sqlRunner{dsn: b.config.Db}
	err = b.runner.init()
	return &b, err
}

func config(fn string) (Config, error) {
	c := Config{}
	dat, err := ioutil.ReadFile(fn)
	if err != nil {
		return c, err
	}

	if err := json.Unmarshal(dat, &c); err != nil {
		return c, err
	}

	return c, nil
}
