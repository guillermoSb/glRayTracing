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
	envMap, _ := gl.NewTexture("parking.bmp")
	r.SetEnvMap(envMap)
	// red, _ := gl.NewColor(1,0,0)
	// green, _ := gl.NewColor(0,1,0)
	semiWhite, _ := gl.NewColor(0, 1, 1)
	// greenMat := gl.NewMaterial(*green, 1, gl.OPAQUE, 1)
	mirrorMat := gl.NewMaterial(*semiWhite, 1, gl.REFLECTIVE, 1.500)
	// redMat := gl.NewMaterial(*red, 1, gl.OPAQUE, 1)

	sphere := gl.NewSphere(numg.V3{0, 0, -15}, 2, *mirrorMat)
	// sphere2 := gl.NewSphere(numg.V3{-5,0,-15}, 1, *greenMat)
	// sphere3 := gl.NewSphere(numg.V3{-2,0,-30}, 4, *redMat)

	// ambientLight1 := gl.NewAmbientLight(gl.White(), 0.2)
	dirLight1 := gl.NewDirLight(numg.V3{-1, 0, -1}, gl.White(), 1)
	ambientLight := gl.NewPointLight(gl.White(), 4, numg.V3{0, 0, -5})
	// Add the objects to the scene
	r.AddToScene(sphere)
	r.AddLightToScene(dirLight1)
	r.AddLightToScene(ambientLight)
	r.GLRender()
	r.GlFinish("output.bmp")
}
