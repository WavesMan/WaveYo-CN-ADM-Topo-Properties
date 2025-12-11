package admtp_test

import (
	"testing"

	admtp "github.com/WavesMan/WaveYo-CN-ADM-Topo-Properties/Open-SDK/Golang-SDK/admtp"
)

func TestDivisionBasics(t *testing.T) {
	d, err := admtp.NewDivision()
	if err != nil {
		t.Fatalf("new division err: %v", err)
	}
	if len(d.ListProvinces()) == 0 {
		t.Fatalf("no provinces")
	}
	if p, ok := d.GetProvinceByID("110000"); !ok || p.Zh == "" {
		t.Fatalf("get province by id failed")
	}
	if p, ok := d.GetProvinceByName("北京市"); !ok || p.ID != "110000" {
		t.Fatalf("get province by name variant failed")
	}
	cs := d.GetCitiesByProvince("130000")
	if len(cs) == 0 {
		t.Fatalf("expect cities for 130000 > 0")
	}
	if c, ok := d.GetCityByName("石家庄"); !ok || c.ID[:4] != "1301" {
		t.Fatalf("get city by name failed")
	}
	ks := d.GetCountiesByCity("130100")
	if len(ks) == 0 {
		t.Fatalf("expect counties for city 130100 > 0")
	}
	if k, ok := d.GetCountyByName("海淀区"); !ok || k.ID[:2] != "11" {
		t.Fatalf("get county by name failed")
	}
}

func TestNegativeLookups(t *testing.T) {
	d, err := admtp.NewDivision()
	if err != nil {
		t.Fatalf("new division err: %v", err)
	}
	if _, ok := d.GetProvinceByID("999999"); ok {
		t.Fatalf("unexpected province")
	}
	if _, ok := d.GetCityByName("不存在城市"); ok {
		t.Fatalf("unexpected city")
	}
	if _, ok := d.GetCountyByID("000000"); ok {
		t.Fatalf("unexpected county")
	}
	r := d.SearchByName("北京")
	if len(r) == 0 {
		t.Fatalf("search no result")
	}
}
