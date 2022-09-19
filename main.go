package main

// turbosquid
import (
	"fmt"
	"guillermoSb/glRayTracing/gl"
	"guillermoSb/glRayTracing/numg"
)

func main() {
	fmt.Println("Hello from RTX program")
	r, _ := gl.NewRenderer(1024, 1024, "")
	red, _ := gl.NewColor(0.8,0.3,0.3)
	gray, _ := gl.NewColor(0.4,0.4,0.4)
	brick := gl.NewMaterial(*red, 100)
	stone := gl.NewMaterial(*gray, 10)

	sphere := gl.NewSphere(numg.V3{0,0,-15}, 2, *brick)
	sphere2 := gl.NewSphere(numg.V3{-3,0,-20}, 1, *stone)
	light1 := gl.NewDirLight(numg.V3{0,-1,-1}, gl.White(), 1)
	ambientLight1 := gl.NewAmbientLight(gl.White(), 0.2)
	// Add the objects to the scene
	r.AddToScene(sphere)
	r.AddToScene(sphere2)
	// Add the lights to the scene
	r.AddLightToScene(light1)
	r.AddLightToScene(ambientLight1)
	r.GLRender()
	r.GlFinish("output.bmp")
}