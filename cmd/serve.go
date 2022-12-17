package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-backend/api"
	"go-backend/conf"
	"log"
)

func serve(cmd *cobra.Command, args []string) {
	gconf, err := conf.LoadGlobalConfig()
	if err != nil {
		log.Println("Failed to load global configuration")
		return
	}

	rs := api.RegisterEndpoints(gconf)
	log.Println("restful server is running")

	endpoint := fmt.Sprintf("%s:%d", gconf.EndPoint.Host, gconf.EndPoint.Port)
	fmt.Println("Listening on", endpoint)
	rs.ListenAndServe(endpoint)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the EssayGenie server",
	Long:  "Start the EssayGenie server",
	Run:   serve,
}
