package main

// turbosquid
import (
	"fmt"
	"guillermoSb/glRayTracing/gl"
	"guillermoSb/glRayTracing/numg"
)

func main() {
	fmt.Println("Hello from RTX program")
	r, _ := gl.NewRenderer(2048, 1024, "hdr.bmp")

	// Env map
	// envMap, _ := gl.NewTexture("parking.bmp")
	// r.SetEnvMap(envMap)
	red, _ := gl.NewColor(1, 0, 0)
	// green, _ := gl.NewColor(0,1,0)
	semiWhite, _ := gl.NewColor(1, 0.77, 0.20)
	mirrorMat := gl.NewMaterial(*semiWhite, 200, gl.OPAQUE, 1.500)
	brickMat := gl.NewMaterial(*red, 200, gl.OPAQUE, 1.500)

	sphere := gl.NewSphere(numg.V3{0, 0, -15}, 2, *mirrorMat)
	sphere2 := gl.NewSphere(numg.V3{5, 0, -10}, 1.4, *brickMat)
	// sphere3 := gl.NewSphere(numg.V3{8, 0, -10}, 1, *brickMat)
	dirLight1 := gl.NewDirLight(numg.V3{-1, 0, -1}, *semiWhite, 0.5)
	ambientLight := gl.NewAmbientLight(gl.White(), 0.2)
	// Add the objects to the scene
	r.AddToScene(sphere)
	r.AddToScene(sphere2)
	// r.AddToScene(sphere3)
	// Add lights to the scene
	r.AddLightToScene(dirLight1)
	r.AddLightToScene(ambientLight)
	r.GLRender()
	r.GlFinish("output.bmp")
}
