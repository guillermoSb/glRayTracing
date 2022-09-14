package main

// turbosquid
import (
	"fmt"
	"guillermoSb/glRayTracing/gl"
)

func main() {
	fmt.Println("Hello from RTX program")
	r, _ := gl.NewRenderer(2040, 1024, "sunshine-bg.bmp")
	r.GLDrawBackground()
	
	
	r.GlFinish("output.bmp")
}