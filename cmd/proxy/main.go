package proxy

import (
	"log"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thaddeuscleo/gopetuah/config"
	"github.com/thaddeuscleo/gopetuah/pkg/proxy"
)

// rootCmd represents the base command when called without any subcommands
var ProxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "A brief proxy of your application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var c config.Config
		errUnmarshall := viper.Unmarshal(&c)

		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		}))
		slog.SetDefault(logger)

		if errUnmarshall != nil {
			log.Panicln(errUnmarshall)
		}

		log.Println("Server start on port 3000")
		server := proxy.New(c)
		err := server.Start()
		if err != nil {
			log.Panicln(err)
		}
	},
}

func init() {

}
