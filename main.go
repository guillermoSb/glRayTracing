package main

// turbosquid
import (
	"fmt"
	"guillermoSb/glRayTracing/gl"
	"guillermoSb/glRayTracing/numg"
)

func main() {
	fmt.Println("Hello from RTX program")
	r, _ := gl.NewRenderer(1024, 1024, "hdr.bmp")

	// Env map
	envMap, _ := gl.NewTexture("parking.bmp")
	r.SetEnvMap(envMap)
	red, _ := gl.NewColor(1,0,0)
	green, _ := gl.NewColor(0,1,0)
	semiWhite, _ := gl.NewColor(0.4,0.4,0.4)
	greenMat := gl.NewMaterial(*green, 1, gl.OPAQUE)
	mirrorMat := gl.NewMaterial(*semiWhite, 1, gl.REFLECTIVE)
	redMat := gl.NewMaterial(*red, 1, gl.OPAQUE)
	

	sphere := gl.NewSphere(numg.V3{0,0,-15}, 2, *mirrorMat)
	sphere2 := gl.NewSphere(numg.V3{-5,0,-15}, 1, *greenMat)
	sphere3 := gl.NewSphere(numg.V3{5,0,-15}, 2, *redMat)
	pointLight1 := gl.NewPointLight(gl.White(),1, numg.V3{0,0,-12})
	ambientLight1 := gl.NewAmbientLight(gl.White(), 0.2)
	// Add the objects to the scene
	r.AddToScene(sphere)
	r.AddToScene(sphere2)
	r.AddToScene(sphere3)
	// r.AddToScene(sphere3)
	// r.AddToScene(sphere2)
	// Add the lights to the scene
	r.AddLightToScene(pointLight1)
	r.AddLightToScene(ambientLight1)
	r.GLRender()
	r.GlFinish("output.bmp")
}