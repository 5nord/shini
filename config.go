package main

import (
	"github.com/pelletier/go-toml"
	"github.com/pelletier/go-toml/query"
	"io"
	"os"
)

type Config struct {
	t *toml.Tree
}

func (c *Config) Get(keys []string) interface{} {
	if len(keys) == 0 {
		return c.t
	}

	var results []interface{}
	for _, v := range keys {
		results = append(results, c.t.Get(v))
	}
	return results
}

func (c *Config) Query(q string) (interface{}, error) {
	result, err := query.CompileAndExecute(q, c.t)
	if err != nil {
		return nil, err
	}

	return result.Values(), err
}

func readConfig(file string) (*Config, error) {
	var r io.Reader
	var err error

	if file == "-" {
		r = os.Stdin
	} else {
		r, err = os.Open(file)
		if err != nil {
			return nil, err
		}
	}

	c := new(Config)
	c.t, err = toml.LoadReader(r)
	if err != nil {
		return nil, err
	}

	return c, nil
}
