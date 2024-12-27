package server

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/common-nighthawk/go-figure"
	"github.com/gin-gonic/gin"
	"github.com/mss-boot-io/mss-boot/core/server"
	"github.com/mss-boot-io/mss-boot/core/server/listener"
	"github.com/spf13/cobra"

	_ "github.com/iniscan-labs/iniscan/apis"
	"github.com/iniscan-labs/iniscan/config"
	"github.com/iniscan-labs/iniscan/router"
)

var (
	StartCmd = &cobra.Command{
		Use:     "server",
		Short:   "start server",
		Long:    "start iniscan server",
		Example: "iniscan server",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func setup() error {
	slog.Info("setup server")

	// 01 setup config
	config.Cfg.Init()

	// 02 setup gin
	r := gin.Default()

	// 03 setup runnable
	runnable := []server.Runnable{
		config.Cfg.Server.Init(
			listener.WithStartedHook(tips),
			listener.WithName("iniscan"),
			listener.WithHandler(r),
		),
	}

	// 04 add runnable
	server.Manage.Add(runnable...)

	router.Init(r.Group("/explorer"))
	return nil
}

func run() error {
	ctx := context.Background()
	return server.Manage.Start(ctx)
}

func tips() {
	figure.NewFigure("iniscan", "rectangles", true).Print()
	fmt.Println()
}
