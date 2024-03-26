// command client receives and displays monitoring parameters
package main

import (
	"fmt"
	"time"

	"github.com/Elyas-Crimean/sysmon/internal/client"
	"github.com/spf13/pflag"
)

func main() {
	host := pflag.StringP("host", "h", "127.0.0.1:8000", "сетевой адрес и порт сервера мониторинга")
	interval := pflag.DurationP("interval", "i", 5*time.Second, "интервал между выдачей значений параметров (1s..30s)")
	window := pflag.DurationP("window", "w", 15*time.Second, "окно для усреднения выборок параметров (5s..2m)")

	if *interval < 1000*time.Millisecond || *interval > 30000*time.Millisecond {
		fmt.Println("недопустимый интервал")
		return
	}
	if *window < 5*time.Second || *window > 120*time.Second {
		fmt.Println("недопустимое окно усреднения")
		return
	}
	client := client.NewClient(*host, *interval, *window)
	if client == nil {
		fmt.Println("ошибка создания gRPC клиента")
		return
	}
}
