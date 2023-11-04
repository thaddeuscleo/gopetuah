package proxy

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/thaddeuscleo/gopetuah/pkg/proxy"
)

// rootCmd represents the base command when called without any subcommands
var ProxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "A brief proxy of your application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Server start on port 3000")
		server := proxy.New(":3000")
		err := server.Start()
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {

}
