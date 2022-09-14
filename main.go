package main

// turbosquid
import (
	"fmt"
	"guillermoSb/glRayTracing/gl"
	"guillermoSb/glRayTracing/numg"
)

func main() {
	fmt.Println("Hello from RTX program")
	r, _ := gl.NewRenderer(256, 256, "")
	sphere := gl.NewSphere(numg.V3{0,0,0}, 1.4)
	r.AddToScene(sphere)
	r.GLRender()
	r.GlFinish("output.bmp")
}