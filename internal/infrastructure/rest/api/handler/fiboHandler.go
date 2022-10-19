package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type getFibonacci struct {
	Sequence []int64 `json: "sequence`
}
func NewGetFibonacci (s []int64) getFibonacci {
	return getFibonacci{
		Sequence: s,
	}
}

func (h *Handler) GetFibonacci(c *gin.Context){
	//get number x
	x, err := strconv.Atoi(c.Query("x"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid x param")
		return
	}
	//get number y
	y, err := strconv.Atoi(c.Query("y"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid y param")
		return
	}
	//get fibonacci sequence
	s, err := h.Services.GetSequence(x, y)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Cannot get sequence")
	}
	c.JSON(http.StatusOK, NewGetFibonacci(s))	
}