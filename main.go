package main

// turbosquid
import (
	"fmt"
	"guillermoSb/glRayTracing/gl"
	"guillermoSb/glRayTracing/numg"
)

func main() {
	fmt.Println("Hello from RTX program")
	r, _ := gl.NewRenderer(1024, 800, "")

	// Env map
	envMap, _ := gl.NewTexture("beach.bmp")
	r.SetEnvMap(envMap)

	brick := gl.NewMaterial(gl.CreateColor(1, 0, 0), 100, gl.OPAQUE, 0, nil)
	grass := gl.NewMaterial(gl.CreateColor(0.2, 1, 0.2), 50, gl.OPAQUE, 0, nil)
	glass := gl.NewMaterial(gl.CreateColor(0, 1, 1), 100, gl.TRANSPARENT, 1.5, nil)
	diamond := gl.NewMaterial(gl.CreateColor(0.9, 0.9, 0.9), 200, gl.TRANSPARENT, 2.4, nil)
	sunMaterial := gl.NewMaterial(gl.CreateColor(1, 1, 0.4), 100, gl.REFLECTIVE, 1.5, nil)
	darkMaterial := gl.NewMaterial(gl.CreateColor(0.4, 0.4, 0.), 10, gl.REFLECTIVE, 1.5, nil)

	sphere1 := gl.NewSphere(numg.V3{-3, 1.5, -10}, 1, *glass)
	sphere2 := gl.NewSphere(numg.V3{0, 1.5, -10}, 1, *diamond)
	sphere3 := gl.NewSphere(numg.V3{3, 1.5, -10}, 1, *brick)

	sphere4 := gl.NewSphere(numg.V3{-3, -1.5, -10}, 1, *sunMaterial)
	sphere5 := gl.NewSphere(numg.V3{0, -1.5, -10}, 1, *darkMaterial)
	sphere6 := gl.NewSphere(numg.V3{3, -1.5, -10}, 1, *grass)

	dirLight1 := gl.NewDirLight(numg.V3{-1, 0, -1}, gl.White(), 0.5)
	ambientLight := gl.NewAmbientLight(gl.White(), 0.2)
	// Add the objects to the scene
	// r.AddToScene(sphere)
	r.AddToScene(sphere1)
	r.AddToScene(sphere2)
	r.AddToScene(sphere3)
	r.AddToScene(sphere4)
	r.AddToScene(sphere5)
	r.AddToScene(sphere6)
	// Add lights to the scene
	r.AddLightToScene(dirLight1)
	r.AddLightToScene(ambientLight)
	r.GLRender()
	r.GlFinish("output.bmp")
}
