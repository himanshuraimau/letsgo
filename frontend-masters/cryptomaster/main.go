package main

import (
	"fmt"
	"log"
	"enghimanshu.tech/go/cryptomaster/api"
	"sync"
)

func GetCurrency(currency string) {
	rate, err := api.GetRate(currency)
	if err ==nil{
		fmt.Printf("1 %s is $%f\n", rate.Currency, rate.Price)
    	} else{
		log.Printf("Error getting rate for %s: %v", currency, err)
}
}

func main(){
	   currencies := []string{"BTC", "ETH", "LTC", "XRP", "SOL"}
       var wg sync.WaitGroup
	   for _, currency := range currencies {
		   wg.Add(1)
		     go func(currencyCode string) {
                GetCurrency(currencyCode)
				wg.Done()	
			 }(currency)
		}
	   wg.Wait()
}