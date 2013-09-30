package main

import (
	"bitbucket.org/liamstask/goose/lib/goose"
	"log"
)

var downCmd = &Command{
	Name:    "down",
	Usage:   "",
	Summary: "Roll back the version by 1",
	Help:    `down extended help here...`,
}

func downRun(cmd *Command, args ...string) {

	conf, err := goose.NewDBConf(*flagPath, *flagEnv)
	if err != nil {
		log.Fatal(err)
	}

	current, err := goose.GetDBVersion(conf)
	if err != nil {
		log.Fatal(err)
	}

	previous, err := goose.GetPreviousDBVersion(conf.MigrationsDir, current)
	if err != nil {
		log.Fatal(err)
	}

	if err = goose.RunMigrations(conf, conf.MigrationsDir, previous); err != nil {
		log.Fatal(err)
	}
}

func init() {
	downCmd.Run = downRun
}
