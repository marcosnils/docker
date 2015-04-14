package main

import (
	"github.com/codegangsta/cli"
	"github.com/docker/libcontainer"
)

var checkpointCommand = cli.Command{
	Name:  "checkpoint",
	Usage: "checkpoint a running container",
	Flags: []cli.Flag{
		cli.StringFlag{Name: "id", Value: "nsinit", Usage: "specify the ID for a container"},
	},
	Action: func(context *cli.Context) {
		container, err := getContainer(context)
		if err != nil {
			fatal(err)
		}
		if err := container.Checkpoint(&libcontainer.CriuOpts{}); err != nil {
			fatal(err)
		}
	},
}