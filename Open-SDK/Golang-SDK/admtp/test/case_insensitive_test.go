package admtp_test

import (
    "testing"
    admtp "github.com/WavesMan/WaveYo-CN-ADM-Topo-Properties/Open-SDK/Golang-SDK/admtp"
)

func TestEnglishCaseInsensitive(t *testing.T) {
    d, err := admtp.NewDivision()
    if err != nil {
        t.Fatalf("new division err: %v", err)
    }
    if p, ok := d.GetProvinceByName("bEiJiNg"); !ok || p.ID != "110000" {
        t.Fatalf("province mixed-case english lookup failed")
    }
    if c, ok := d.GetCityByName("sHiJiAzHuAnG"); !ok || c.ID[:4] != "1301" {
        t.Fatalf("city mixed-case english lookup failed")
    }
    if k, ok := d.GetCountyByName("hAiDiAnQu"); !ok || k.ID != "110108" {
        t.Fatalf("county mixed-case english lookup failed")
    }
}
