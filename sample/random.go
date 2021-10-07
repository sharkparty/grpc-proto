package sample

import (
	"grpc-proto/pb"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBoolean() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "ATI")
}

func randomLaptopBrand() string {
	return randomStringFromSet("Dell", "Apple", "Lenovo")
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomCPUName(b string) string {
	if b == "Intel" {
		return randomStringFromSet(
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
		)
	}
	return randomStringFromSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500U",
		"Ryzen 3 PRO 3200GE",
	)
}

func randomGPUName(b string) string {
	if b == "NVIDIA" {
		return randomStringFromSet(
			"RTX-2060",
			"RTX-2070",
			"GTX 1660-Ti",
		)
	}
	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
	)
}

func randomLaptopName(b string) string {
	switch b {
	case "Apple":
		return randomStringFromSet(
			"Macbook Pro 15",
			"Macbook Pro 16",
			"Macbook Pro 13",
		)
	case "Dell":
		return randomStringFromSet(
			"Latitude",
			"XPS",
			"Alienware",
		)
	default:
		return randomStringFromSet(
			"Thinkpad X1",
			"Thinkpad P1",
			"Thinkpad P53",
		)
	}
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9

	resolution := &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}

	return resolution
}

func randomUUID() string {
	return uuid.New().String()
}
