package main
import(
	"fmt"
	"os"
	"net/http"
	"encoding/json"
	"net/url"
	"strings"
	"bufio"
)
type Item struct{
	Text	string	`json:text`
}
func InputPrompt(label string) string {
    var s string
    r := bufio.NewReader(os.Stdin)
    for {
        fmt.Fprint(os.Stderr, label+" ")
        s, _ = r.ReadString('\n')
        if s != "" {
            break
        }
    }
    return strings.TrimSpace(s)
}
func main(){
        message := InputPrompt("> ")
	text :=  url.QueryEscape(message)
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
