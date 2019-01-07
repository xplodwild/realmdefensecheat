package realmdefense_test

import (
	"reflect"
	"testing"

	"github.com/xplodwild/realmdefensecheat/realmdefense"
)

func TestGetValueFromOSV(t *testing.T) {
	tests := []struct {
		name string
		args realmdefense.SaveDataOsvVal
		want int
	}{
		{name: "Value 1", args: realmdefense.SaveDataOsvVal{O: -84967, S: 1, V: -84965}, want: 2},
		{name: "Value 2", args: realmdefense.SaveDataOsvVal{O: 55849, S: 1, V: 55849}, want: 0},
		{name: "Value 3", args: realmdefense.SaveDataOsvVal{O: 24539, S: -1, V: 24537}, want: 2},
		{name: "Value 4", args: realmdefense.SaveDataOsvVal{O: 97423, S: 1, V: 97425}, want: 2},
		{name: "Value 5", args: realmdefense.SaveDataOsvVal{O: 51741, S: -1, V: 51711}, want: 30},
		{name: "Value 6", args: realmdefense.SaveDataOsvVal{O: -83817, S: -1, V: -83819}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := realmdefense.GetValueFromOSV(tt.args); got != tt.want {
				t.Errorf("GetValueFromOSV() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValueFromIVEntry(t *testing.T) {
	tests := []struct {
		name string
		args realmdefense.SaveDataIvEntry
		want int
	}{
		{name: "value 1", args: realmdefense.SaveDataIvEntry{
			G: realmdefense.SaveDataOsvVal{O: -30560, S: -1, V: -31649},
			U: realmdefense.SaveDataOsvVal{O: -68871, S: 1, V: -67785},
		}, want: 3},
		{name: "value 2", args: realmdefense.SaveDataIvEntry{
			G: realmdefense.SaveDataOsvVal{O: -34278, S: 1, V: -33007},
			U: realmdefense.SaveDataOsvVal{O: -37186, S: -1, V: -38318},
		}, want: 139},
		{name: "value 3", args: realmdefense.SaveDataIvEntry{
			G: realmdefense.SaveDataOsvVal{O: -39297, S: -1, V: -1746872},
			U: realmdefense.SaveDataOsvVal{O: -40842, S: -1, V: -1734342},
		}, want: 14075},
		{name: "value 4", args: realmdefense.SaveDataIvEntry{
			G: realmdefense.SaveDataOsvVal{O: -87403, S: 1, V: 82424},
			U: realmdefense.SaveDataOsvVal{O: -29434, S: -1, V: -186101},
		}, want: 13160},
		{name: "value 5", args: realmdefense.SaveDataIvEntry{
			G: realmdefense.SaveDataOsvVal{O: -77710, S: 1, V: -75813},
			U: realmdefense.SaveDataOsvVal{O: -90466, S: -1, V: -91592},
		}, want: 771},
		{name: "value 6", args: realmdefense.SaveDataIvEntry{
			G: realmdefense.SaveDataOsvVal{O: -20894, S: 1, V: -17449},
			U: realmdefense.SaveDataOsvVal{O: 60415, S: 1, V: 60415},
		}, want: 3445},
		{name: "value 7", args: realmdefense.SaveDataIvEntry{
			G: realmdefense.SaveDataOsvVal{O: -63420, S: 1, V: -63370},
			U: realmdefense.SaveDataOsvVal{O: -58676, S: 1, V: -58676},
		}, want: 50},
		{name: "value 8", args: realmdefense.SaveDataIvEntry{
			G: realmdefense.SaveDataOsvVal{O: 37146, S: 1, V: 37266},
			U: realmdefense.SaveDataOsvVal{O: -55572, S: 1, V: -55572},
		}, want: 120},
		{name: "value 9", args: realmdefense.SaveDataIvEntry{
			G: realmdefense.SaveDataOsvVal{O: 22857, S: 1, V: 23357},
			U: realmdefense.SaveDataOsvVal{O: 23424, S: 1, V: 23424},
		}, want: 500},
		{name: "value 10", args: realmdefense.SaveDataIvEntry{
			G: realmdefense.SaveDataOsvVal{O: 55316, S: 1, V: 56316},
			U: realmdefense.SaveDataOsvVal{O: -46448, S: 1, V: -46448},
		}, want: 1000},
		{name: "value 11", args: realmdefense.SaveDataIvEntry{
			G: realmdefense.SaveDataOsvVal{O: -83464, S: -1, V: -83964},
			U: realmdefense.SaveDataOsvVal{O: -46448, S: 1, V: -46448},
		}, want: 500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := realmdefense.GetValueFromIVEntry(tt.args); got != tt.want {
				t.Errorf("GetValueFromIVEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeOSVValue(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{name: "value 1", args: 100, want: 100},
		{name: "value 2", args: -100, want: -100},
		{name: "value 3", args: 0, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := realmdefense.MakeOSVValue(tt.args); !reflect.DeepEqual(realmdefense.GetValueFromOSV(got), tt.want) {
				t.Errorf("MakeOSVValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
