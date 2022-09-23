package gl

import (
	"guillermoSb/glRayTracing/numg"
	"log"
	"math"
	"os"
)

type renderer struct {
	width, height               uint
	pixels                      [][]color
	scene                       []figure
	lights                      []light
	background                  texture
	camPosition                 numg.V3
	nearPlane, aspectRatio, fov float64
	currentColor                color
	envMap                      *texture
}

const MAX_RECURSION_DEPTH = 2

// Creates a new renderer and sends the reference
// - width: width of the renderer
// - height: height of the renderer
// - background: name of the background
func NewRenderer(width, height uint, background string) (*renderer, error) {
	r := renderer{width: width, height: height, pixels: [][]color{}}
	// Create the pixel array
	for x := 0; x < int(width); x++ {
		col := []color{}
		for y := 0; y < int(height); y++ {
			clearClr, _ := NewColor(0.6, 0.6, 0.6) // Use black as the background color
			col = append(col, *clearClr)           // append the color to the column
		}
		r.pixels = append(r.pixels, col) // append the column to the pixels
	}
	// Create an empty array of figures
	r.scene = []figure{}
	// Camera position
	r.camPosition = numg.V3{X: 0, Y: 0, Z: 0}
	// NearPlane
	r.nearPlane = 0.1
	// AspectRatio
	r.aspectRatio = float64(width) / float64(height)
	// FOV
	r.fov = 60
	// Load the background
	if background != "" {
		t, errT := NewTexture(background)
		if errT != nil {
			return nil, errT
		}
		r.background = *t
	}
	// Default color is white
	r.currentColor = color{1, 1, 1}
	// Return the renderer
	return &r, nil
}

func (r *renderer) SetEnvMap(envMap *texture) {
	r.envMap = envMap
}

// Change current color
func (r *renderer) ChangeColor(clr color) {
	r.currentColor = clr
}

// Get the pixel array
func (r *renderer) Pixels() [][]color {
	return r.pixels
}

// Append an item to the scene
func (r *renderer) AddToScene(figure figure) {
	r.scene = append(r.scene, figure)
}

func (r *renderer) AddLightToScene(light light) {
	r.lights = append(r.lights, light)
}

// Draw the background
func (r *renderer) GLDrawBackground() {
	for x := 0; x < int(r.width); x++ {
		for y := 0; y < int(r.height); y++ {
			ux := (float32(x) / float32(r.width))  // value for x 0 -1
			uy := (float32(y) / float32(r.height)) // value for y 0 -1
			texX := int(ux * float32(r.background.width))
			texY := int(uy * float32(r.background.height))
			r.GLPoint(numg.V2{X: float64(x), Y: float64(y)}, r.background.pixels[texY][texX])
		}
	}
}

// Does the render of the scene with Ray Tracing
func (r *renderer) GLRender() {
	// Proyecion
	t := math.Tan((r.fov*math.Pi)/360) * r.nearPlane
	ri := t * r.aspectRatio
	for x := 0; x < int(r.width); x++ {
		for y := 0; y < int(r.height); y++ {
			// Create the NDC Coordinates
			px := ((float64(x) / float64(r.width)) * 2) - 1.0
			py := ((float64(y) / float64(r.height)) * 2) - 1.0

			// ? Por que se multiplica
			px *= ri
			py *= t
			// Direction of the ray normalized
			// ? Por que la direccion tiene en z - el nearPlane
			direction := numg.NormalizeV3(numg.V3{X: px, Y: py, Z: -r.nearPlane})
			// Cast a ray on that direction
			rayColor := r.GLCastRay(r.camPosition, direction, nil, 0)
			if rayColor != nil {
				r.GLPoint(numg.V2{X: float64(x), Y: float64(y)}, *rayColor)
			}
		}
	}
}

func (r *renderer) sceneIntersect(origin, direction numg.V3, sceneObject *struct{ material }) *intersect {
	var intersect *intersect = nil
	depth := math.Inf(1)
	for _, object := range r.scene {
		hit := object.rayIntersect(origin, direction)
		if hit != nil {
			if &(hit.obj) != sceneObject {
				if sceneObject != nil && hit.obj == *sceneObject {
					continue
				}
				if hit.distance < depth {
					intersect = hit
					depth = hit.distance
				}
			}
		}
	}
	return intersect
}

func (r *renderer) GLCastRay(origin, direction numg.V3, sceneObject *struct{ material }, recursion int) *color {
	var sceneIntersect *intersect = r.sceneIntersect(origin, direction, sceneObject)
	if sceneIntersect != nil && recursion < MAX_RECURSION_DEPTH {
		finalColor := Black()
		objectColor := sceneIntersect.obj.material.diffuse
		if sceneIntersect.obj.material.matType == OPAQUE {
			for _, light := range r.lights {
				lightColor := light.getLightColor(r.camPosition, *sceneIntersect)
				shadowIntensity := 0.0
				finalColor.r += lightColor.r
				finalColor.g += lightColor.g
				finalColor.b += lightColor.b
				var shadowIntersect *intersect = nil
				if light.getLightType() == DIR_TYPE {
					shadowIntersect = r.sceneIntersect(sceneIntersect.point, numg.MultiplyVectorWithConstant(*light.getDirection(), -1), &sceneIntersect.obj)
				} else if light.getLightType() == POINT_TYPE {
					lightDir := numg.Subtract(*light.getOrigin(), sceneIntersect.point)
					lightDir = numg.NormalizeV3(lightDir)
					shadowIntersect = r.sceneIntersect(sceneIntersect.point, numg.MultiplyVectorWithConstant(lightDir, -1), &sceneIntersect.obj)
				}
				if shadowIntersect != nil {
					shadowIntensity = 1
				}
				finalColor.r *= 1 - shadowIntensity
				finalColor.g *= 1 - shadowIntensity
				finalColor.b *= 1 - shadowIntensity
			}

			finalColor.r *= objectColor.r
			finalColor.g *= objectColor.g
			finalColor.b *= objectColor.b
		} else if sceneIntersect.obj.material.matType == REFLECTIVE {
			// reflection vector
			reflection := numg.ReflectionVector(numg.MultiplyVectorWithConstant(direction, -1.0), numg.MultiplyVectorWithConstant(sceneIntersect.normal, 1))
			reflectColor := r.GLCastRay(sceneIntersect.point, reflection, &sceneIntersect.obj, recursion+1)
			if reflectColor != nil {
				finalColor.r = reflectColor.r * objectColor.r
				finalColor.g = reflectColor.g * objectColor.g
				finalColor.b = reflectColor.b * objectColor.b
			} else {
				finalColor = objectColor
			}
			for _, light := range r.lights {
				lightColor := light.getLightColor(r.camPosition, *sceneIntersect)
				shadowIntensity := 0.0
				finalColor.r += lightColor.r
				finalColor.g += lightColor.g
				finalColor.b += lightColor.b
				var shadowIntersect *intersect = nil
				if light.getLightType() == DIR_TYPE {
					shadowIntersect = r.sceneIntersect(sceneIntersect.point, numg.MultiplyVectorWithConstant(*light.getDirection(), -1), &sceneIntersect.obj)
				} else if light.getLightType() == POINT_TYPE {
					lightDir := numg.Subtract(*light.getOrigin(), sceneIntersect.point)
					lightDir = numg.NormalizeV3(lightDir)
					shadowIntersect = r.sceneIntersect(sceneIntersect.point, numg.MultiplyVectorWithConstant(lightDir, -1), &sceneIntersect.obj)
				}
				if shadowIntersect != nil {
					shadowIntensity = 1
				}
				finalColor.r *= 1 - shadowIntensity
				finalColor.g *= 1 - shadowIntensity
				finalColor.b *= 1 - shadowIntensity
			}
		} else if sceneIntersect.obj.material.matType == TRANSPARENT {
			// Detect if the ray comes from outside of the object or inside
			outside := (numg.V3DotProduct(direction, sceneIntersect.normal)) < 0
			biasVector := numg.MultiplyVectorWithConstant(sceneIntersect.normal, 0.001)
			reflectionVector := numg.ReflectionVector(direction, sceneIntersect.normal)
			originVector := sceneIntersect.point
			if outside {
				originVector = numg.Add(originVector, biasVector)
			} else {
				originVector = numg.Subtract(originVector, biasVector)
			}
			reflectColor := r.GLCastRay(originVector, reflectionVector, nil, recursion+1)
			refractColor := Black()
			kr := numg.Fresnel(sceneIntersect.normal, direction, sceneIntersect.obj.material.ior)
			if kr < 1 {
				refractVector := numg.Refract(sceneIntersect.normal, direction, sceneIntersect.obj.material.ior)
				refractOrigin := sceneIntersect.point
				if outside {
					refractOrigin = numg.Subtract(refractOrigin, biasVector)
				} else {
					refractOrigin = numg.Add(refractOrigin, biasVector)
				}
				refractColorResult := r.GLCastRay(refractOrigin, refractVector, nil, recursion+1)
				refractColor.r = refractColorResult.r
				refractColor.g = refractColorResult.g
				refractColor.b = refractColorResult.b
			}
			finalColor.r += (reflectColor.r*kr + refractColor.r*(1-kr))
			finalColor.g += (reflectColor.g*kr + refractColor.g*(1-kr))
			finalColor.b += (reflectColor.b*kr + refractColor.b*(1-kr))
		}

		finalColor.r = math.Min(1, finalColor.r)
		finalColor.g = math.Min(1, finalColor.g)
		finalColor.b = math.Min(1, finalColor.b)

		return &(finalColor)
	} else if r.envMap != nil {
		return r.envMap.GetEnvColor(direction)
	}
	return nil
}

// Draw a point on the screen
func (r *renderer) GLPoint(point numg.V2, clr color) {
	// Check tht the point is within the screen bounds
	if int(point.X) < 0 ||
		int(point.X) >= int(r.width) ||
		int(point.Y) < 0 ||
		int(point.Y) >= int(r.height) {
		return
	}
	// Set the color on the pixel
	r.pixels[int(point.X)][int(point.Y)] = clr
}

// Create the renderer with the pixel array
// filename: The name of the file
func (r *renderer) GlFinish(fileName string) {
	// Attempt to open the file
	f, err := os.Create(fileName)

	// Check if the file was successfully created
	if err != nil {
		log.Fatal(err)
	}
	// Example, writing 5 x 5 image
	defer f.Close() // Close the file when the process is done
	f.Write([]byte("B"))
	f.Write([]byte("M"))
	f.Write(numg.Dword(uint32(r.width) * uint32(r.height*3))) // File Size
	f.Write([]byte{0, 0})                                     // Reserved
	f.Write([]byte{0, 0})                                     // Reserved
	f.Write([]byte{54, 0, 0, 0})                              // ?
	f.Write([]byte{40, 0, 0, 0})                              // Header Size
	f.Write(numg.Dword(uint32(r.width)))                      // Width
	f.Write(numg.Dword(uint32(r.height)))                     // Height
	f.Write([]byte{1, 0})                                     // Plane
	f.Write([]byte{24, 0})                                    // BPP
	f.Write([]byte{0, 0, 0, 0})
	f.Write([]byte{0, 0, 0, 0})
	f.Write([]byte{0, 0, 0, 0})
	f.Write([]byte{0, 0, 0, 0})
	f.Write([]byte{0, 0, 0, 0})
	f.Write([]byte{0, 0, 0, 0})
	// Pixel Data
	for i := 0; i < int(r.height); i++ {
		for j := 0; j < int(r.width); j++ {
			f.Write(r.pixels[j][i].Bytes())
		}
	}
}

// func (r * renderer) GLViewMatrix(translate, rotate numg.V3) {
// 	camMatrix := glCreateObjectMatrix(translate, rotate, V3{1,1,1})
// 	viewMatrix, err := numg.InverseOfMatrix(camMatrix)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	r.viewMatrix = viewMatrix

// }
// TODO: Be able to load images and use any aspect ratio
