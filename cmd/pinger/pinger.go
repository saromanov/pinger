package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/saromanov/pinger/internal/server"
	"github.com/saromanov/pinger/config"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)


func run(c *config.Config) error {
	app, err := app.New(c)
	if err != nil {
		return fmt.Errorf("unable to init app: %v", err)
	}
	setupServer(app, c)
	return nil
}

// parseConfig provides parsing of the config .yml file
func parseConfig(path string) (*config.Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %v", err)
	}
	var c *config.Config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, fmt.Errorf("unable to parse .config.yml: %v", err)
	}

	return c, nil
}

func main() {
	app := cli.NewApp()
	app.Name = "pinger"
	app.Usage = "checking of availability of sites"
	app.Commands = []cli.Command{
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "path to .yml config",
			Action: func(c *cli.Context) error {
				configPath := c.Args().First()
				config, err := parseConfig(configPath)
				if err != nil {
					panic(err)
				}
				if err := run(config); err != nil {
					panic(err)
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}