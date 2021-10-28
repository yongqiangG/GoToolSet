package main

import (
	"fmt"
	"github.com/GoProject/GoToolSet/internal/json"
	"testing"
)

func TestJsonPretty(t *testing.T) {
	fmt.Printf("%s", json.JSONPretty("{\"id\":1,\"name\":\"poloxue\"}", true))
}
