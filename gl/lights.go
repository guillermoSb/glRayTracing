package gl

import "guillermoSb/glRayTracing/numg"


type light interface {
	getIntensity() float64
	getColor() color
	getLightType() int
	getDirection() *numg.V3
}


type dirLight struct {
	direction numg.V3
	color color
	intensity float64
	lightType int
}

const DIR_TYPE = 0
const AMBIENT_TYPE = 1

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



type ambientLight struct {
	color color
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