package cmd

import (
	"net/http"
	"log"
	"os"
	"fmt"

	"my_cache/server"
	"my_cache/cache"
	"github.com/spf13/cobra"
	_ "github.com/joho/godotenv/autoload"
)

// listenerCmd represents the listener command
var listenerCmd = &cobra.Command{
	Use:   "listener",
	Short: "Start HTTP server",
	Long: "Starts HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		redisClient := cache.NewRedisClient()
		server      := server.NewServer(redisClient)

		port := os.Getenv("PORT")
		fmt.Println("Running on port", port)

		if err := http.ListenAndServe(":"+port, server); err != nil {
			log.Fatalf("could not listen on port %s %v", port, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listenerCmd)
}
