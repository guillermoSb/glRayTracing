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
	// r.GLDrawBackground()
	// Env map
	envMap, _ := gl.NewTexture("nebula.bmp")
	r.SetEnvMap(envMap)

	brick := gl.NewMaterial(gl.CreateColor(1, 1, 0), 20, gl.OPAQUE, 1.1, nil)

	torus := gl.NewTorus(1, 2, *brick)
	// sphere := gl.NewSphere(numg.V3{0, 0, -2}, 0.5, *brick)

	// Add the objects to the scene
	// r.AddToScene(sphere)
	dirLight := gl.NewDirLight(numg.V3{0, 1, -1}, gl.White(), 1)
	ambientLight := gl.NewAmbientLight(gl.White(), 0.2)
	r.AddToScene(torus)
	r.AddLightToScene(ambientLight)
	r.AddLightToScene(dirLight)

	r.GLRender()
	r.GlFinish("output.bmp")
}
