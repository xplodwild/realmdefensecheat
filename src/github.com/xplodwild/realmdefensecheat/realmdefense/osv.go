package realmdefense

import "math/rand"

func GetValueFromOSV(val SaveDataOsvVal) int {
	// val.O *= val.S
	// val.V *= val.S
	return (val.V - val.O) * val.S
}

func GetValueFromIVEntry(val SaveDataIvEntry) int {
	return GetValueFromOSV(val.G) - GetValueFromOSV(val.U)
}

func MakeOSVValue(value int) SaveDataOsvVal {
	// Generate a random S value, either 1 or -1
	s := rand.Intn(2)
	if s == 0 {
		s = -1
	}

	// Generate a random O value, multiplied by S since it will be remultiplied on the other end
	o := rand.Intn(100000) * s

	return SaveDataOsvVal{
		O: o,
		S: s,
		V: o + value,
	}
}

func MakeIVEntryValue(value int) SaveDataIvEntry {
	return SaveDataIvEntry{
		G: MakeOSVValue(value),
		U: MakeOSVValue(0),
	}
}
