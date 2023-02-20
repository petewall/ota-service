package cmd

import (
	"fmt"
	. "github.com/petewall/ota-service/v2/internal"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "ota-service",
	Short: "A brief description of your application",
	RunE: func(cmd *cobra.Command, args []string) error {
		api := &API{
			Updater: &UpdaterImpl{
				DeviceService: &DeviceServiceImpl{
					Host:       viper.GetString("device_service.host"),
					Port:       viper.GetInt("device_service.port"),
					HTTPClient: http.DefaultClient,
				},
				FirmwareService: &FirmwareServiceImpl{
					Host: viper.GetString("firmware_service.host"),
					Port: viper.GetInt("firmware_service.port"),
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

	rootCmd.Flags().String("device-service-host", "", "Device service host")
	_ = viper.BindPFlag("device_service.host", rootCmd.Flags().Lookup("port"))
	_ = viper.BindEnv("device_service.host", "DEVICE_SERVICE_HOST")

	rootCmd.Flags().Int("device-service-port", 5000, "Device service port")
	_ = viper.BindPFlag("device_service.port", rootCmd.Flags().Lookup("db.port"))
	_ = viper.BindEnv("device_service.port", "DEVICE_SERVICE_PORT")

	rootCmd.Flags().String("firmware-service-host", "", "Firmware service host")
	_ = viper.BindPFlag("firmware_service.host", rootCmd.Flags().Lookup("port"))
	_ = viper.BindEnv("firmware_service.host", "FIRMWARE_SERVICE_HOST")

	rootCmd.Flags().Int("firmware-service-port", 5050, "Firmware service port")
	_ = viper.BindPFlag("firmware_service.port", rootCmd.Flags().Lookup("db.port"))
	_ = viper.BindEnv("firmware_service.port", "FIRMWARE_SERVICE_HOST")
}
