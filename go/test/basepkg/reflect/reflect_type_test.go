package basePkgT

import (
	"reflect"
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

func TestRefType(t *testing.T) {
	type Time time.Time

	now := Time(time.Now())

	v := reflect.ValueOf(now)
	assert.Equal(t, v.Kind(), reflect.Struct)

	assert.Equal(t, v.Type().Name(), "Time")
	assert.Equal(t, v.Type().String(), "basePkgT.Time")

	assert.Equal(t, v.Type().String(), "basePkgT.Time")
}
