package cmd

import (
	"fmt"
	"net/http"
	"os"

	deviceService "github.com/petewall/device-service/lib"
	firmwareService "github.com/petewall/firmware-service/lib"
	. "github.com/petewall/ota-service/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "ota-service",
	Short: "A brief description of your application",
	RunE: func(cmd *cobra.Command, args []string) error {
		api := &API{
			Updater: &UpdaterImpl{
				DeviceService: &deviceService.Client{
					URL: viper.GetString("device_service"),
				},
				FirmwareService: &firmwareService.Client{
					URL: viper.GetString("firmware_service"),
				},
			},
			LogOutput: cmd.OutOrStdout(),
		}

		port := viper.GetInt("port")
		cmd.Printf("Listening on port %d\n", port)
		return http.ListenAndServe(fmt.Sprintf(":%d", port), api.GetMux())
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Int("port", 8266, "Port to listen on")
	_ = viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))
	_ = viper.BindEnv("port", "PORT")

	rootCmd.Flags().String("device-service", "", "Device service host")
	_ = viper.BindPFlag("device_service", rootCmd.Flags().Lookup("device-service"))
	_ = viper.BindEnv("device_service", "DEVICE_SERVICE")

	rootCmd.Flags().String("firmware-service", "", "Firmware service host")
	_ = viper.BindPFlag("firmware_service", rootCmd.Flags().Lookup("firmware-service"))
	_ = viper.BindEnv("firmware_service", "FIRMWARE_SERVICE")

	rootCmd.SetOut(rootCmd.OutOrStdout())
}
