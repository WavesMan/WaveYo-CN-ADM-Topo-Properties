package admtp

import "testing"

func TestCoverageForAdmtpPackage(t *testing.T) {
    d, err := NewDivision()
    if err != nil {
        t.Fatalf("new division err: %v", err)
    }
    if len(d.ListProvinces()) == 0 || len(d.ListCounties()) == 0 {
        t.Fatalf("lists should be non-empty")
    }
    if p, ok := d.GetProvinceByName("北京市"); !ok || p.ID != "110000" {
        t.Fatalf("province variant lookup failed")
    }
    if p, ok := d.GetProvinceByName("Beijing"); !ok || p.ID != "110000" {
        t.Fatalf("province english lookup failed")
    }
    if c, ok := d.GetCityByName("石家庄"); !ok || c.ID[:4] != "1301" {
        t.Fatalf("city lookup failed")
    }
    if k, ok := d.GetCountyByName("海淀区"); !ok || k.ID != "110108" {
        t.Fatalf("county lookup failed")
    }
    if len(d.GetCitiesByProvince("110000")) != 0 {
        t.Fatalf("municipality cities should be empty")
    }
    if len(d.GetCountiesByCity("110100")) == 0 {
        t.Fatalf("beijing counties by city prefix should be non-empty")
    }
    r := d.SearchByName("北京")
    if len(r) == 0 {
        t.Fatalf("search should return results")
    }
    if _, ok := d.GetProvinceByID("999999"); ok {
        t.Fatalf("unexpected province exists")
    }
}
