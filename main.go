package main

// turbosquid
import (
	"fmt"
	"guillermoSb/glRayTracing/gl"
	"guillermoSb/glRayTracing/numg"
)

func main() {
	fmt.Println("Hello from RTX program")
	r, _ := gl.NewRenderer(400, 400, "")
	// r.GLDrawBackground()
	// Env map
	envMap, _ := gl.NewTexture("nebula.bmp")
	r.SetEnvMap(envMap)

	brick := gl.NewMaterial(gl.CreateColor(1, 0, 0), 20, gl.OPAQUE, 0, nil)
	grass := gl.NewMaterial(gl.CreateColor(0.2, 1, 0.2), 1, gl.OPAQUE, 1, nil)
	ceil := gl.NewMaterial(gl.CreateColor(0.7, 0.7, 0.7), 1, gl.OPAQUE, 1, nil)
	glass := gl.NewMaterial(gl.CreateColor(0, 1, 1), 100, gl.TRANSPARENT, 1.5, nil)
	diamond := gl.NewMaterial(gl.CreateColor(0.9, 0.9, 0.9), 200, gl.TRANSPARENT, 2.4, nil)
	// sunMaterial := gl.NewMaterial(gl.CreateColor(1, 1, 0.4), 100, gl.REFLECTIVE, 1.5, nil)
	// darkMaterial := gl.NewMaterial(gl.CreateColor(0.4, 0.4, 0.), 10, gl.REFLECTIVE, 1.5, nil)
	plane1 := gl.NewPlane(*grass, numg.V3{0, -10, -10}, numg.V3{0, 1, 0})
	plane2 := gl.NewPlane(*ceil, numg.V3{0, 10, 0}, numg.V3{0, -1, 0})
	plane3 := gl.NewPlane(*brick, numg.V3{-20, 0, 0}, numg.V3{1, 0, 0})
	plane4 := gl.NewPlane(*brick, numg.V3{20, 0, 0}, numg.V3{-1, 0, 0})

	backgroundBox := gl.NewAABB(*brick, numg.V3{0, 0, -40}, numg.V3{20, 10, 1})

	box1 := gl.NewAABB(*glass, numg.V3{-2, 0, -10}, numg.V3{2, 2, 2})
	box2 := gl.NewAABB(*diamond, numg.V3{2, 0, -10}, numg.V3{2, 2, 2})

	dirLight := gl.NewDirLight(numg.V3{0, 0, -1}, gl.CreateColor(0.2, 0.2, 1), 1)
	dirLight2 := gl.NewDirLight(numg.V3{-1, 0, -1}, gl.CreateColor(1, 1, 1), 0.5)
	dirLight3 := gl.NewDirLight(numg.V3{1, 0, -1}, gl.CreateColor(1, 1, 1), 0.5)
	ambientLight := gl.NewAmbientLight(gl.CreateColor(0.2, 0.2, 1), 0.6)
	// Add the objects to the scene
	// r.AddToScene(sphere)
	r.AddToScene(plane1)
	r.AddToScene(plane2)
	r.AddToScene(plane3)
	r.AddToScene(plane4)
	r.AddToScene(backgroundBox)

	r.AddToScene(box1)
	r.AddToScene(box2)

	// Add lights to the scene
	r.AddLightToScene(dirLight)
	r.AddLightToScene(dirLight2)
	r.AddLightToScene(dirLight3)
	r.AddLightToScene(ambientLight)

	r.GLRender()
	r.GlFinish("output.bmp")
}
