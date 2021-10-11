package main

import (
	"fmt"

	"github.com/krastomer/shoptree/backend/pkg/repository/mariadb"
)

// func init() {
// 	viper.AutomaticEnv()

// 	switch viper.GetString("env") {
// 	case "dev":
// 		viper.SetDefault("address", "127.0.0.1")
// 	default:
// 		// gin.SetMode("release")
// 		viper.SetDefault("address", "0.0.0.0")
// 	}

// 	if viper.GetString("port") == "" {
// 		viper.SetDefault("port", "8080")
// 	}

// }
const (
	dsn     = "root:password@tcp(127.0.0.1:3306)/shoptree?charset=utf8&parseTime=True&loc=Local"
	address = "127.0.0.1:8080"
)

func main() {
	// address := fmt.Sprintf("%s:%s", viper.GetString("address"), viper.GetString("port"))
	repo := mariadb.NewRepository(dsn)

	cust, err := repo.GetCustomer(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(cust)

	// handler := rest.NewHandler()

	// server := &http.Server{
	// 	Addr:         address,
	// 	Handler:      handler,
	// 	ReadTimeout:  10 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }

	// server.ListenAndServe()

}
