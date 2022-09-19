package gl

import (
	"guillermoSb/glRayTracing/numg"
	"log"
	"math"
	"os"
)
type renderer struct {
	width, height uint
	pixels [][]color
	scene []figure
	lights []light
	background texture
	camPosition numg.V3
	nearPlane, aspectRatio,fov float64
	currentColor color

}

// Creates a new renderer and sends the reference
// - width: width of the renderer
// - height: height of the renderer
// - background: name of the background
func NewRenderer(width, height uint, background string)(*renderer, error) {
	r := renderer{width: width, height: height, pixels: [][]color{}} 
	// Create the pixel array
	for x := 0; x < int(width); x++ {
		col := []color{}
		for y := 0; y < int(height); y++ {
			clearClr, _ := NewColor(0,0,0)	// Use black as the background color
			col = append(col, *clearClr)	// append the color to the column
		}
		r.pixels = append(r.pixels, col)	// append the column to the pixels
	}
	// Create an empty array of figures
	r.scene = []figure{}
	// Camera position
	r.camPosition = numg.V3{X: 0,Y: 0,Z: 0}
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
	r.currentColor = color{1,1,1}
	// Return the renderer
	return &r, nil
}

// Change current color
func(r *renderer) ChangeColor(clr color) {
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
			ux := (float32(x)/float32(r.width))	// value for x 0 -1
			uy := (float32(y)/float32(r.height))	// value for y 0 -1
			texX := int(ux * float32(r.background.width))	
			texY := int(uy * float32(r.background.height))
			r.GLPoint(numg.V2{X:float64(x),Y:float64(y)}, r.background.pixels[texY][texX])
		}	
	}
}

// Does the render of the scene with Ray Tracing
func (r *renderer) GLRender() {
	// Proyecion
	t := math.Tan((r.fov * math.Pi)/360) * r.nearPlane
	ri := t * r.aspectRatio
	for x := 0; x < int(r.width); x++ {
		for y := 0; y < int(r.height); y++ {
			// Create the NDC Coordinates
			px := ((float64(x)/float64(r.width)) * 2) - 1.0
			py := ((float64(y)/float64(r.height)) * 2) - 1.0
			
			// ? Por que se multiplica
			px *= ri
			py *= t
			// Direction of the ray normalized
			// ? Por que la direccion tiene en z - el nearPlane
			direction := numg.NormalizeV3(numg.V3{X: px, Y: py,Z: -r.nearPlane})
			// Cast a ray on that direction
			rayColor := r.GLCastRay(r.camPosition, direction)	
			if rayColor != nil {
				r.GLPoint(numg.V2{X: float64(x),Y: float64(y)}, *rayColor)
			}
		}
	}
}

func (r *renderer) sceneIntersect(origin, direction numg.V3, sceneObject *struct{material}) *intersect {
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

func (r *renderer) GLCastRay(origin, direction numg.V3) *color {
	var intersect *intersect = r.sceneIntersect(origin, direction, nil)
	if intersect != nil {
		finalColor := Black()
		objectColor := intersect.obj.material.diffuse

		for _, light := range r.lights {
			if light.getLightType() == DIR_TYPE {
				dirLight := numg.MultiplyVectorWithConstant(*light.getDirection(), -1)
				intensity := numg.V3DotProduct(dirLight, numg.NormalizeV3(intersect.normal))
				intensity = math.Max(0, intensity)
				diffuseColor, _ := NewColor(light.getColor().r * intensity, light.getColor().g * intensity, light.getColor().b * intensity)
				// Shadows
				shadowIntensity := 0.0
				shadowIntersect := r.sceneIntersect(intersect.point, dirLight, &intersect.obj)
				if shadowIntersect != nil {
					shadowIntensity = 1
				}
				// Specularity
			
				rS := numg.ReflectionVector(dirLight, intersect.normal)

				viewDir := numg.NormalizeV3(numg.Subtract(r.camPosition, intersect.point))
				specIntensity := numg.V3DotProduct(rS, viewDir)
				specIntensity = math.Max(0, specIntensity)
				specIntensity = math.Pow(specIntensity, intersect.obj.material.specularity)
				specColor, _ := NewColor(light.getColor().r * specIntensity, light.getColor().g * specIntensity, light.getColor().b * specIntensity)
				
				diffuseColor.r += specColor.r
				diffuseColor.g += specColor.g
				diffuseColor.b += specColor.b

				finalColor.r += diffuseColor.r * light.getIntensity()
				finalColor.g += diffuseColor.g * light.getIntensity()
				finalColor.b += diffuseColor.b * light.getIntensity()

				finalColor.r *= 1 - shadowIntensity
				finalColor.g *= 1 - shadowIntensity
				finalColor.b *= 1 - shadowIntensity

			} else if light.getLightType() == AMBIENT_TYPE {
				ambientLightColor := light.getColor()
				ambientLightColor.r = ambientLightColor.r * light.getIntensity()
				ambientLightColor.g = ambientLightColor.g * light.getIntensity()
				ambientLightColor.b = ambientLightColor.b * light.getIntensity()

				finalColor.r += ambientLightColor.r
				finalColor.g += ambientLightColor.g
				finalColor.b += ambientLightColor.b
			} else if light.getLightType() == POINT_TYPE {
				// Calculate the light direction
				lightDir := numg.Subtract(*light.getOrigin(), intersect.point)
				lightDir = numg.NormalizeV3(lightDir)
				// Calculate the intensity of the light
				intensity := numg.V3DotProduct(numg.NormalizeV3(intersect.point), lightDir)
				intensity = math.Max(0, intensity)
				// Calculate the color
				diffuseColor, _ := NewColor(light.getColor().r * intensity, light.getColor().g * intensity, light.getColor().b * intensity)
				// Shadows
				shadowIntensity := 0.0
				shadowIntersect := r.sceneIntersect(intersect.point, lightDir, &intersect.obj)
				if shadowIntersect != nil {
					shadowIntensity = 1
				}
				// Specularity
			
				rS := numg.ReflectionVector(lightDir, intersect.normal)

				viewDir := numg.NormalizeV3(numg.Subtract(r.camPosition, intersect.point))
				specIntensity := numg.V3DotProduct(rS, viewDir)
				specIntensity = math.Max(0, specIntensity)
				specIntensity = math.Pow(specIntensity, intersect.obj.material.specularity)
				specColor, _ := NewColor(light.getColor().r * specIntensity, light.getColor().g * specIntensity, light.getColor().b * specIntensity)
				
				diffuseColor.r += specColor.r
				diffuseColor.g += specColor.g
				diffuseColor.b += specColor.b
				
				// Attenuation
				distanceVector := numg.Subtract(intersect.point, *light.getOrigin())
				
				d := math.Abs(distanceVector.Magnitude())
				a := 1.0	// Constant attenuation
				b := 1.0	// Linear attenuation
				c := 0.0	// Quadratic attenuatio

				attenuation := 1.0 / (a+b*d+c*math.Pow(d,2))

				finalColor.r += diffuseColor.r * light.getIntensity() * attenuation
				finalColor.g += diffuseColor.g * light.getIntensity() * attenuation
				finalColor.b += diffuseColor.b * light.getIntensity() * attenuation

				finalColor.r *= 1 - shadowIntensity
				finalColor.g *= 1 - shadowIntensity
				finalColor.b *= 1 - shadowIntensity
			}
		}

		finalColor.r *= objectColor.r 
		finalColor.g *= objectColor.g
		finalColor.b *= objectColor.b

		finalColor.r = math.Min(1, finalColor.r)
		finalColor.g = math.Min(1, finalColor.g)
		finalColor.b = math.Min(1, finalColor.b)
		 
		return &(finalColor)
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
	defer f.Close()	// Close the file when the process is done
	f.Write([]byte("B"))
	f.Write([]byte("M"))
	f.Write(numg.Dword(uint32(r.width) * uint32(r.height * 3)))	// File Size
	f.Write([]byte{0, 0})	// Reserved
	f.Write([]byte{0, 0})	// Reserved
	f.Write([]byte{54, 0, 0, 0 })	// ?
	f.Write([]byte{40, 0, 0, 0})	// Header Size
	f.Write(numg.Dword(uint32(r.width)))		// Width
	f.Write(numg.Dword(uint32(r.height)))		// Height
	f.Write([]byte{1, 0})	// Plane
	f.Write([]byte{24, 0})	// BPP
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
	f.Write([]byte{0,0,0,0})
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
