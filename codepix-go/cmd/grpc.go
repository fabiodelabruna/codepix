package cmd

import (
	"os"

	"github.com/fabiodelabruna/codepix/codepix-go/application/grpc"
	"github.com/fabiodelabruna/codepix/codepix-go/infrastructure/db"
	"github.com/spf13/cobra"
)

var portNumber int

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start a gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		grpc.StartGrpcServer(database, portNumber)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
	grpcCmd.Flags().IntVarP(&portNumber, "port", "p", 50051, "gRPC Port Number")
}
