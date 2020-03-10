// THIS FILE IS AUTO-GENERATED
package service

import (
	"hc/characteristic"
)

const TypeLightbulb = "43"

type Lightbulb struct {
	*Service

	On *characteristic.On
}

func NewLightbulb() *Lightbulb {
	svc := Lightbulb{}
	svc.Service = New(TypeLightbulb)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	return &svc
}
