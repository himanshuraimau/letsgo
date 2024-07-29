package api

import (
	"fmt"
	"net/http"
	"strings"

	"enghimanshu.tech/go/cryptomaster/datatypes"
)
const apiURL = "https://api.cex.io/api/tickt/%s/USD"

func GetRate(currency string) (datatypes.Rate,error)  {
	upCurrency:= strings.ToUpper(currency)
	res,err := http.GET(fmt.Sprintf(apiURL,upCurrency))
}