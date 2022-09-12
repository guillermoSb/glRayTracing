package main

// turbosquid
import (
	"fmt"
	"guillermoSb/glRayTracing/gl"
	"guillermoSb/glRayTracing/numg"
)

func main() {
	fmt.Println("Hello from RTX program")
	r, _ := gl.NewRenderer(20, 20)
	b, _ := gl.NewColor(0,0,1)
	r.GLPoint(numg.V2{X: 2,Y: 5}, *b)
	r.GlFinish("out.bmp")
}