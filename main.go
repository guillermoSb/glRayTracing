package main

// turbosquid
import (
	"fmt"
	"guillermoSb/glRayTracing/gl"
)

func main() {
	fmt.Println("Hello from RTX program")
	r, _ := gl.NewRenderer(400, 400, "")
	r.GLRender()
	r.GlFinish("output.bmp")
}
