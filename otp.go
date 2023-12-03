package otp

import (
	"context"
	"time"

	"github.com/cdvelop/token"
)

// newOTP crea y agrega un nuevo OTP al mapa
func (o *Otp) newOTP() otp {
	n := otp{
		key:     token.BuildUniqueKey(32),
		created: time.Now(),
	}

	o.rm[n.key] = n
	return n
}

// verifyOTP se asegurará de que exista un OTP
// y devolverá true en caso afirmativo
// También eliminará la clave para que no se pueda reutilizar
func (o *Otp) verifyOTP(otp string) bool {
	// Verify OTP is existing
	if _, ok := o.rm[otp]; !ok {
		// otp does not exist
		return false
	}
	delete(o.rm, otp)
	return true
}

// esta función se asegurará periódicamente que se eliminen los "OTP" antiguos
// que han expirado según el período de retención especificado.
// La función utiliza un ticker para realizar la retención en intervalos regulares
// y se detendrá si se recibe una señal de cancelación desde el contexto.
// es bloqueante, así que ejecútelo como una goroutine
func (o *Otp) retention(ctx context.Context, retentionPeriod time.Duration) {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			for _, otp := range o.rm {
				// comprobar si ha caducado
				if otp.created.Add(retentionPeriod).Before(time.Now()) {
					delete(o.rm, otp.key)
				}
			}
		case <-ctx.Done():
			return

		}
	}
}
