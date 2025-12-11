package admtp_test

import (
	"testing"

	admtp "github.com/WavesMan/WaveYo-CN-ADM-Topo-Properties/Open-SDK/Golang-SDK/admtp"
)

func TestEnglishNameLookups(t *testing.T) {
	d, err := admtp.NewDivision()
	if err != nil {
		t.Fatalf("new division err: %v", err)
	}
	if p, ok := d.GetProvinceByName("Beijing"); !ok || p.ID != "110000" {
		t.Fatalf("province english lookup failed")
	}
	if c, ok := d.GetCityByName("Shijiazhuang"); !ok || c.ID[:4] != "1301" {
		t.Fatalf("city english lookup failed")
	}
	if k, ok := d.GetCountyByName("Haidianqu"); !ok || k.ID != "110108" {
		t.Fatalf("county english lookup failed")
	}
}

func TestProvinceNameVariants(t *testing.T) {
	d, err := admtp.NewDivision()
	if err != nil {
		t.Fatalf("new division err: %v", err)
	}
	if p, ok := d.GetProvinceByName("内蒙古自治区"); !ok || p.ID != "150000" {
		t.Fatalf("Inner Mongolia variant failed")
	}
	if p, ok := d.GetProvinceByName("宁夏省"); !ok || p.ID != "640000" {
		t.Fatalf("Ningxia implicit suffix failed")
	}
}

func TestListsAndEmptyCases(t *testing.T) {
	d, err := admtp.NewDivision()
	if err != nil {
		t.Fatalf("new division err: %v", err)
	}
	if len(d.ListProvinces()) == 0 || len(d.ListCities()) == 0 || len(d.ListCounties()) == 0 {
		t.Fatalf("list functions should return non-empty")
	}
	if len(d.GetCitiesByProvince("110000")) != 0 { // municipality has empty cities file
		t.Fatalf("expect no cities for 110000")
	}
}

func TestSearchEnglish(t *testing.T) {
	d, err := admtp.NewDivision()
	if err != nil {
		t.Fatalf("new division err: %v", err)
	}
	r := d.SearchByName("Beij")
	if len(r) == 0 {
		t.Fatalf("search english partial failed")
	}
}
