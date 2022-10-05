package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	fibologic "github.com/ninja-dark/fibonacci_testtask/internal/fiboLogic"
)

type Router struct{
	*http.ServeMux
	fibo *fibologic.Fibo
}

func NewRouter(fibo *fibologic.Fibo) *Router{
	r := &Router{
		ServeMux: http.NewServeMux(),
		fibo: fibo,
	}

	r.Handle("/fibonacci", http.HandlerFunc(r.GetSequence ))
	return r
}

type Fibo struct{
	Sequence []int `json:"sequence"`
}

func (rt *Router) GetSequence(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet{
		http.Error(w, "Bad method", http.StatusMethodNotAllowed)
		return
	}
	f := r.URL.Query().Get("num")
	if f == ""{
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	n, _ := strconv.Atoi(f)

	fibo, err := rt.fibo.GetSequence(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	_ = json.NewEncoder(w).Encode(
		Fibo{
			Sequence: fibo,
		},
	) 
}