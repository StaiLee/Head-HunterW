package changeme

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetFromApi(_URI string, target interface{}) error {
	res, err := http.Get(_URI)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	err2 := json.NewDecoder(res.Body).Decode(target)
	if err2 == nil {
		return json.NewDecoder(res.Body).Decode(target)
	}
	fmt.Println(err2)
	return nil
}
