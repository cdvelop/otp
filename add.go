package otp

func Add() (o *Otp) {

	o = &Otp{
		rm: make(map[string]otp),
	}

	return o
}
