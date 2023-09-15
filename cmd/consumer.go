package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"example.com/go-inmem-loki/api"
	"example.com/go-inmem-loki/constant"
	"example.com/go-inmem-loki/consumer"
	"example.com/go-inmem-loki/db"
	"example.com/go-inmem-loki/kafka"
	"example.com/go-inmem-loki/util"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(consumerCmd)
}

var consumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "ranker - feature-store - consumer",
	Long:  `ranker - feature-store - consumer`,
	Run: func(cmd *cobra.Command, _ []string) {
		go func() {
			fmt.Println(http.ListenAndServe("0.0.0.0:6060", nil))
		}()
		mainCtx, stop := signal.NotifyContext(cmd.Context(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		log.Info().Msgf("ENVIRONMENT: %s", util.GetConfig().GoEnv)

		if err := kafka.CreateTopicIfNotExists(kafka.RankerFeatureStoreUpdatesTopic, 1, 1); err != nil {
			panic(err)
		}

		db.Init()

		var wg sync.WaitGroup
		wg.Add(1)
		go kafka.EventListener(mainCtx, consumer.HandleUpdate, &wg)
		t := time.AfterFunc(constant.ExpiryInterval, consumer.ExpireFeatures)

		wg.Add(1)
		go api.StartConsumerServer(mainCtx, &wg)

		// Listen for signals
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigs
		log.Info().Str("Component", "serverCommand").Msgf("Got signal, beginning shutdown %s", sig)
		wg.Wait()
		kafka.Close()
		t.Stop()
		log.Info().Str("Component", "serverCommand").Msg("All goroutines finished, shutting down...!")
	},
}
