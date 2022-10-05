package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	fibologic "github.com/ninja-dark/fibonacci_testtask/internal/fiboLogic"
	"github.com/ninja-dark/fibonacci_testtask/internal/infrastructure/rest/api/router"
)


func main(){
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	fibo := fibologic.NewFibo()
	h := router.NewRouter(fibo)



}