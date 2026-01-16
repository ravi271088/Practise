// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func main() {
	//	g := gin.Default()
	//	g.GET("/hello", getHello)
	//	g.Run(":8080")
	r := mux.NewRouter()
	r.HandleFunc("/data/{aa}", getData).Methods("GET")
	http.ListenAndServe(":8080", r)

}

func getHello(c *gin.Context) {
	name := c.DefaultQuery("name", "world")
	c.JSON(404, gin.H{
		"message": "hello " + name,
	})
}

func getData(w http.ResponseWriter, r *http.Request) {
	//d := mux.Vars(r)
	//response := d["aa"]
	//fmt.Println(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]string{"message": "New resource created"}
	json.NewEncoder(w).Encode(response)
}
