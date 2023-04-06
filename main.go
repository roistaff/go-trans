package main
import(
	"fmt"
	"os"
	"net/http"
	"encoding/json"
	"net/url"
)
type Item struct{
	Text	string	`json:text`
}

func main(){
	args := os.Args
	text :=  url.QueryEscape(args[1])
	url := "https://script.google.com/macros/s/AKfycbwq5WhR2nY9xv1rIK6LE_Gk1KbTuRpqfdUa4Z_hZL6VymipGMB0-dsmg3KWyN66gtXZTQ/exec?text=" + text + "&sourse=en&target=ja"
	resp,err := http.Get(url)
	if err != nil {
                panic(err)
        }
        defer resp.Body.Close()
	var data Item
        err = json.NewDecoder(resp.Body).Decode(&data)
        if err != nil {
         panic(err)
         }
	 fmt.Println(data.Text)
}
