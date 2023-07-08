package productcontroller

import (
	"net/http"

	"github.com/gslayer0/go-middle/helpers"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"id":   1,
			"nama": "Laptop Acer",
			"stok": 1000,
		},
		{
			"id":   2,
			"nama": "Laptop Apple",
			"stok": 1000,
		},
		{
			"id":   3,
			"nama": "Laptop Asus",
			"stok": 1000,
		},
	}

	helpers.ResponseJSON(w, http.StatusOK, data)
}
