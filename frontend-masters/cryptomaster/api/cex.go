package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"enghimanshu.tech/go/cryptomaster/datatypes"
)
const apiURL = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate,error)  {
	if len(currency) !=3{
		return nil,fmt.Errorf("Invalid currency code %s",currency)
	}
	upCurrency:= strings.ToUpper(currency)
	res,err := http.Get(fmt.Sprintf(apiURL,upCurrency))
	if err !=nil{
			return nil,err
	} 
	var response CEXResponse
	if res.StatusCode == http.StatusOK{
		bodyBytes,err:=io.ReadAll(res.Body)
		if err !=nil{
			return nil,err
		}
		
		err =json.Unmarshal(bodyBytes,&response)
		if err !=nil{
			return nil,err
		}
	} else{
		return nil,fmt.Errorf("API returned status code %d",res.StatusCode)
	}
	rate := datatypes.Rate{Currency: currency,Price:response.Bid}
	return &rate,nil 

}