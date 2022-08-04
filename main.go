package main

// This is the most minimal blinky example and should run almost everywhere.

import (
	"golitter/bts7960"
	"machine"
	"time"
)

/*
int R_IS = 1;
int R_EN = 2;
int R_PWM = 3;
int L_IS = 4;
int L_EN = 5;
int L_PWM = 6;
*/

const (
	rEn = machine.D2
	lEn = machine.D3

	rPwm = machine.PD5
	lPwm = machine.PD6

	maxSpeed = 200
)

func blink(led machine.Pin) {
	led.Low()
	time.Sleep(time.Millisecond * 500)

	led.High()
	time.Sleep(time.Millisecond * 500)

}

func main() {
	var period uint64 = 0

	pwm := machine.Timer0

	if err := pwm.Configure(machine.PWMConfig{Period: period}); err != nil {
		println(err.Error())
		return
	}

	println("starting...")

	bts := bts7960.New(lEn, rEn, lPwm, rPwm, pwm)
	err := bts.Configure()
	if err != nil {
		println("bts configure: " + err.Error())
		return
	}

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	bts.Enable()
	count := 0
	for {
		if count > 5 {
			bts.Stop()
			bts.Disable()
			println("spinning done")
			continue
		}
		println("main loop")

		bts.Right(60)
		blink(led)
		count++

		// for i := 0; i < maxSpeed; i++ {
		// 	println("spin: " + strconv.Atoi)
		// }
	}
}
