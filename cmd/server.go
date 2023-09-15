package cmd

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"example.com/go-inmem-loki/api"
	"example.com/go-inmem-loki/grpc"
	"example.com/go-inmem-loki/inmemory"
	"example.com/go-inmem-loki/kafka"
	"example.com/go-inmem-loki/monitoring"
	"example.com/go-inmem-loki/populate"
	"example.com/go-inmem-loki/util"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "ranker - feature-store",
	Long:  `ranker - feature-store`,
	Run: func(cmd *cobra.Command, _ []string) {
		go func() {
			fmt.Println(http.ListenAndServe("0.0.0.0:6060", nil))
		}()
		monitoring.Init()
		mainCtx, stop := signal.NotifyContext(cmd.Context(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		log.Info().Msgf("ENVIRONMENT: %s", util.GetConfig().GoEnv)

		var wg sync.WaitGroup
		wg.Add(1)
		m := inmemory.Init(mainCtx, &wg)

		if util.GetConfig().GoEnv != "test" {
			populate.Populate(mainCtx, m)
		}

		wg.Add(1)
		go api.StartServer(mainCtx, m, &wg)

		wg.Add(1)
		go grpc.StartListener(mainCtx, m, &wg)

		// Listen for signals
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigs
		api.STATUS.Store(0) // Start to fail readiness probes

		log.Info().Str("Component", "serverCommand").Msgf("Got signal, beginning shutdown %s", sig)
		wg.Wait()
		kafka.Close()
		log.Info().Str("Component", "serverCommand").Msg("All goroutines finished, shutting down...!")
	},
}
