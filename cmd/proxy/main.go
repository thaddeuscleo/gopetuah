package proxy

import (
	"log"

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
		if errUnmarshall != nil {
			log.Println(errUnmarshall)
		}

		log.Println("Server start on port 3000")
		server := proxy.New(c)
		err := server.Start()
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {

}
