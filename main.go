package main

import (
	"fmt"
	"github.com/juancavalpso/urbansolutions/configu"
	"github.com/juancavalpso/urbansolutions/modelo"
	"net/http"
)

func main() {
	fmt.Println("hola")
	strF := paquet.M
	fmt.Println(strF)
	fmt.Println(paquet.Mensaje("import tt 1"))

	http.HandleFunc("/materiales", get)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func get(w http.ResponseWriter, req *http.Request) {
	//entity := modelo.MyEntity{}
	//w.Write(entity.Data)

	entity := modelo.GetOutil()
	w.Write(entity.Data)
}
