package main

// turbosquid
import (
	"fmt"
	"guillermoSb/glRayTracing/gl"
	"guillermoSb/glRayTracing/numg"
	"math/rand"
)

func main() {
	fmt.Println("Hello from RTX program")
	r, _ := gl.NewRenderer(2048, 1024, "")
	// Env map
	envMap, _ := gl.NewTexture("nebula.bmp")
	r.SetEnvMap(envMap)

	spherePoints := []numg.V3{
		{-5, -3, -2},
		{-4, -3, -2.6},
		{-3.3, -2.5, -2.7},
		{-2.7, -2, -2.8},
		{-2.5, -2, -2.8},
		{-2, -1.5, -3},

		{-4, -2, -1},
		{-4, -1, -2.6},
		{-3.3, -1.5, -2.7},
		{-2.7, -1, -2.8},
		{-2.5, -1, -2.8},
		{-2, -0.5, -3},
	}

	gold := gl.NewMaterial(gl.CreateColor(1, 1, 0), 20, gl.REFLECTIVE, 1.1, nil)
	water := gl.NewMaterial(gl.CreateColor(0.8, 0.8, 1), 50, gl.OPAQUE, 1.1, nil)
	black := gl.NewMaterial(gl.Black(), 2, gl.TRANSPARENT, 1.1, nil)
	fire := gl.NewMaterial(gl.CreateColor(221.0/255.0, 87.0/255.0, 28.0/255.0), 2, gl.TRANSPARENT, 2, nil)
	fire2 := gl.NewMaterial(gl.CreateColor(253.0/255.0, 161.0/255.0, 114.0/255.0), 2, gl.TRANSPARENT, 2, nil)

	torus := gl.NewTorus(0.5, 1.5, *gold)
	// box := gl.NewAABB(*gold, numg.V3{-5, 0, 0}, numg.V3{1, 1, 1})
	sphere := gl.NewSphere(numg.V3{0, 0, -2}, 1.3, *black)

	planetEarth := gl.NewSphere(numg.V3{-8, -4, -2}, 3, *water)

	// Add the objects to the scene

	// Black Hole
	r.AddToScene(sphere)
	r.AddToScene(torus)

	// r.AddToScene(box)
	// r.AddToScene(sphere)
	// Planet Earth
	r.AddToScene(planetEarth)

	// Add debree objects
	for _, spherePoint := range spherePoints {
		debreeCount := rand.Intn(2) + 1
		for i := 0; i < debreeCount; i++ {
			point2 := spherePoint
			point2.Y += 0.4
			point2.Z -= 0.4
			point3 := spherePoint
			point2.Y -= 0.2
			point2.Z -= 0.1

			if rand.Float64() > 0.5 {
				sphere := gl.NewSphere(spherePoint, (rand.Float64()*(0.5) + 0.1), *fire)
				sphere1 := gl.NewSphere(point2, (rand.Float64()*(0.5) + 0.1), *fire)
				sphere2 := gl.NewSphere(point3, (rand.Float64()*(0.5) + 0.1), *fire)
				r.AddToScene(sphere)
				r.AddToScene(sphere1)
				r.AddToScene(sphere2)
			} else {
				sphere := gl.NewSphere(spherePoint, (rand.Float64()*(0.5) + 0.1), *fire2)
				sphere1 := gl.NewSphere(point2, (rand.Float64()*(0.5) + 0.1), *fire2)
				sphere2 := gl.NewSphere(point3, (rand.Float64()*(0.5) + 0.1), *fire2)
				r.AddToScene(sphere)
				r.AddToScene(sphere1)
				r.AddToScene(sphere2)
			}
		}

	}

	// Lights
	dirLight := gl.NewDirLight(numg.V3{0, 1, -1}, gl.White(), 1)
	ambientLight := gl.NewAmbientLight(gl.White(), 0.2)

	pointLight := gl.NewPointLight(gl.CreateColor(1, 0, 0), 1, numg.V3{-5, -4, 0})
	pointLight2 := gl.NewPointLight(gl.CreateColor(1, 0, 0), 1, numg.V3{-4, -4, -3})

	r.AddLightToScene(ambientLight)
	r.AddLightToScene(dirLight)

	// Fire on the planet
	r.AddLightToScene(pointLight)
	r.AddLightToScene(pointLight2)

	r.GLRender()
	r.GlFinish("output.bmp")
}
