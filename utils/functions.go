package changeme

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//notre fonction qui permet de manipuler le json de notre API
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

//fonction qui va nous etre utile pour manipuler les slices dans notre barre de recherche pour plus tard
func RemoveArtist(slice []Artist, s int) []Artist {
	return append(slice[:s], slice[s+1:]...)
}
