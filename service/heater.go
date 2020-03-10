package service

import (
	"github.com/MrEhbr/hc/characteristic"
)

type Heater struct {
	*HeaterCooler

	HeatingThresholdTemperature *characteristic.HeatingThresholdTemperature
}

func NewHeater() *Heater {
	svc := Heater{}

	svc.HeatingThresholdTemperature = characteristic.NewHeatingThresholdTemperature()
	svc.AddCharacteristic(svc.HeatingThresholdTemperature.Characteristic)

	return &svc
}
