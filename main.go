package main

// turbosquid
import (
	"fmt"
	"guillermoSb/glRayTracing/gl"
)

func main() {
	fmt.Println("Hello from RTX program")
	r, _ := gl.NewRenderer(10, 10)
}