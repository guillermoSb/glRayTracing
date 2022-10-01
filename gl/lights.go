package gl

import (
	"guillermoSb/glRayTracing/numg"
	"math"
)

type light interface {
	getIntensity() float64
	getColor() color
	getLightType() int
	getDirection() *numg.V3
	getOrigin() *numg.V3
	getLightColor(camPosition numg.V3, intersect intersect) *color
}

type dirLight struct {
	direction numg.V3
	color     color
	intensity float64
	lightType int
}

const DIR_TYPE = 0
const AMBIENT_TYPE = 1
const POINT_TYPE = 3

// Creates a new Directional Light object
func NewDirLight(direction numg.V3, color color, intensity float64) *dirLight {
	dL := dirLight{direction: numg.NormalizeV3(direction), color: color, intensity: intensity, lightType: DIR_TYPE}
	return &dL
}

func (l *dirLight) getIntensity() float64 {
	return l.intensity
}

func (l *dirLight) getColor() color {
	return l.color
}

func (l *dirLight) getLightType() int {
	return l.lightType
}

func (l *dirLight) getDirection() *numg.V3 {
	return &l.direction
}

func (l *dirLight) getOrigin() *numg.V3 {
	return nil
}

func getSpecColor(dirLight, camPosition numg.V3, intersect intersect, lightColor color) color {
	rS := numg.ReflectionVector(dirLight, intersect.normal)

	viewDir := numg.NormalizeV3(numg.Subtract(camPosition, intersect.point))
	specIntensity := numg.V3DotProduct(rS, viewDir)
	specIntensity = math.Max(0, specIntensity)
	specIntensity = math.Pow(specIntensity, intersect.obj.material.specularity)
	specColor, _ := NewColor(lightColor.r*specIntensity, lightColor.g*specIntensity, lightColor.b*specIntensity)

	diffuseColor := Black()
	diffuseColor.r += specColor.r
	diffuseColor.g += specColor.g
	diffuseColor.b += specColor.b
	return diffuseColor
}

func (light *dirLight) getLightColor(camPosition numg.V3, intersect intersect) *color {
	dirLight := numg.MultiplyVectorWithConstant(*light.getDirection(), -1)
	intensity := numg.V3DotProduct(dirLight, numg.NormalizeV3(intersect.normal))
	intensity = math.Max(0, intensity)
	diffuseColor, _ := NewColor(light.getColor().r*intensity, light.getColor().g*intensity, light.getColor().b*intensity)
	// Specularity
	specColor := getSpecColor(dirLight, camPosition, intersect, light.getColor())
	diffuseColor.r += specColor.r
	diffuseColor.g += specColor.g
	diffuseColor.b += specColor.b
	lightColor := Black()
	lightColor.r += diffuseColor.r * light.getIntensity()
	lightColor.g += diffuseColor.g * light.getIntensity()
	lightColor.b += diffuseColor.b * light.getIntensity()
	return &lightColor
}

type ambientLight struct {
	color     color
	intensity float64
	lightType int
}

// Creates a new Ambient Light object
func NewAmbientLight(color color, intensity float64) *ambientLight {
	aL := ambientLight{color: color, intensity: intensity, lightType: AMBIENT_TYPE}
	return &aL
}

func (l *ambientLight) getIntensity() float64 {
	return l.intensity
}

func (l *ambientLight) getColor() color {
	return l.color
}

func (l *ambientLight) getLightType() int {
	return l.lightType
}

func (l *ambientLight) getDirection() *numg.V3 {
	return nil
}

func (l *ambientLight) getOrigin() *numg.V3 {
	return nil
}

func (light *ambientLight) getLightColor(camPosition numg.V3, intersect intersect) *color {
	ambientLightColor := light.getColor()
	ambientLightColor.r = ambientLightColor.r * light.getIntensity()
	ambientLightColor.g = ambientLightColor.g * light.getIntensity()
	ambientLightColor.b = ambientLightColor.b * light.getIntensity()
	return &ambientLightColor
}

type pointLight struct {
	color     color
	intensity float64
	lightType int
	origin    numg.V3
}

// Creates a new Point Light object
func NewPointLight(color color, intensity float64, origin numg.V3) *pointLight {
	pL := pointLight{color: color, intensity: intensity, lightType: POINT_TYPE, origin: origin}
	return &pL
}

func (l *pointLight) getIntensity() float64 {
	return l.intensity
}

func (l *pointLight) getColor() color {
	return l.color
}

func (l *pointLight) getLightType() int {
	return l.lightType
}

func (l *pointLight) getDirection() *numg.V3 {
	return nil
}
func (l *pointLight) getOrigin() *numg.V3 {
	return &l.origin
}

func (light *pointLight) getLightColor(camPosition numg.V3, intersect intersect) *color {
	// Calculate the light direction
	lightDir := numg.Subtract(*light.getOrigin(), intersect.point)
	lightDir = numg.NormalizeV3(lightDir)
	// Calculate the intensity of the light
	intensity := numg.V3DotProduct(numg.NormalizeV3(intersect.point), lightDir)
	intensity = math.Max(0, intensity)
	// Calculate the color
	diffuseColor, _ := NewColor(light.getColor().r*intensity, light.getColor().g*intensity, light.getColor().b*intensity)

	// Specularity

	rS := numg.ReflectionVector(lightDir, intersect.normal)

	viewDir := numg.NormalizeV3(numg.Subtract(camPosition, intersect.point))
	specIntensity := numg.V3DotProduct(rS, viewDir)
	specIntensity = math.Max(0, specIntensity)
	specIntensity = math.Pow(specIntensity, intersect.obj.material.specularity)
	specColor, _ := NewColor(light.getColor().r*specIntensity, light.getColor().g*specIntensity, light.getColor().b*specIntensity)

	diffuseColor.r += specColor.r
	diffuseColor.g += specColor.g
	diffuseColor.b += specColor.b

	// Attenuation
	distanceVector := numg.Subtract(intersect.point, *light.getOrigin())

	d := math.Abs(distanceVector.Magnitude())
	a := 1.0 // Constant attenuation
	b := 1.0 // Linear attenuation
	c := 0.2 // Quadratic attenuatio

	attenuation := 1.0 / (a + b*d + c*math.Pow(d, 2))
	lightColor := Black()
	lightColor.r += diffuseColor.r * light.getIntensity() * attenuation
	lightColor.g += diffuseColor.g * light.getIntensity() * attenuation
	lightColor.b += diffuseColor.b * light.getIntensity() * attenuation
	return &lightColor
}
