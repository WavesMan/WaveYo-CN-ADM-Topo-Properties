package admtp_test

import (
	"testing"

	admtp "github.com/WavesMan/WaveYo-CN-ADM-Topo-Properties/Open-SDK/Golang-SDK/admtp"
)

func TestNameIndexesViaAPI(t *testing.T) {
	d, err := admtp.NewDivision()
	if err != nil {
		t.Fatalf("new division err: %v", err)
	}
	if p, ok := d.GetProvinceByName("北京市"); !ok || p.ID != "110000" {
		t.Fatalf("variant name mapping invalid")
	}
	if p, ok := d.GetProvinceByName("Beijing"); !ok || p.ID != "110000" {
		t.Fatalf("english name mapping invalid")
	}
	if c, ok := d.GetCityByName("Shijiazhuang"); !ok || c.ID[:4] != "1301" {
		t.Fatalf("city english index invalid")
	}
	if k, ok := d.GetCountyByName("Haidianqu"); !ok || k.ID != "110108" {
		t.Fatalf("county english index invalid")
	}
}
