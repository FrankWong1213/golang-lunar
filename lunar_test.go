package lunar

import (
	"testing"
	"time"
)

//go test -v
func TestAverage(t *testing.T) {
	println(Lunar(time.Now().Format("20060102")))
}
