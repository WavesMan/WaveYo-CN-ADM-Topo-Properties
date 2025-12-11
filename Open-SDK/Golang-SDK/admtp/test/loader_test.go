package admtp_test

import (
	"testing"

	admtp "github.com/WavesMan/WaveYo-CN-ADM-Topo-Properties/Open-SDK/Golang-SDK/admtp"
)

func TestDataLoadsViaPublicAPI(t *testing.T) {
	d, err := admtp.NewDivision()
	if err != nil {
		t.Fatalf("new division err: %v", err)
	}
	if len(d.ListProvinces()) == 0 {
		t.Fatalf("no provinces loaded")
	}
	if p, ok := d.GetProvinceByID("110000"); !ok || p.Zh != "北京" {
		t.Fatalf("beijing not found via API")
	}
	cs := d.GetCitiesByProvince("130000")
	if len(cs) == 0 {
		t.Fatalf("expect cities for 130000 > 0")
	}
	ks := d.GetCountiesByCity("110100")
	if len(ks) == 0 {
		t.Fatalf("expect counties for Beijing city prefix > 0")
	}
}
