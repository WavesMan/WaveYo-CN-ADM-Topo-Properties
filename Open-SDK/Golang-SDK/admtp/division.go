package admtp

import "strings"

type Division struct {
    provinces []Province
    cities []City
    counties []County
    provinceByID map[string]Province
    cityByID map[string]City
    countyByID map[string]County
    provinceNameIndex map[string]string
    cityNameIndex map[string]string
    countyNameIndex map[string]string
}

func NewDivision() (*Division, error) {
    prov, err := loadProvinces()
    if err != nil {
        return nil, err
    }
    d := &Division{
        provinces: prov,
        provinceByID: make(map[string]Province),
        cityByID: make(map[string]City),
        countyByID: make(map[string]County),
    }
    for _, p := range prov {
        d.provinceByID[p.ID] = p
    }
    for _, p := range prov {
        if cs, err := loadCities(p.ID); err == nil {
            d.cities = append(d.cities, cs...)
            for _, c := range cs {
                d.cityByID[c.ID] = c
            }
        }
        if ks, err := loadCounties(p.ID); err == nil {
            d.counties = append(d.counties, ks...)
            for _, k := range ks {
                d.countyByID[k.ID] = k
            }
        }
    }
    d.provinceNameIndex = buildProvinceNameIndex(d.provinces)
    d.cityNameIndex = buildNameIndex(wrapCity(d.cities))
    d.countyNameIndex = buildNameIndex(wrapCounty(d.counties))
    return d, nil
}

func wrapCity(items []City) []nameableCity {
    r := make([]nameableCity, 0, len(items))
    for _, it := range items {
        r = append(r, nameableCity{it})
    }
    return r
}

func wrapCounty(items []County) []nameableCounty {
    r := make([]nameableCounty, 0, len(items))
    for _, it := range items {
        r = append(r, nameableCounty{it})
    }
    return r
}

type nameableCity struct{ City }
type nameableCounty struct{ County }

func (c nameableCity) GetID() string { return c.ID }
func (c nameableCity) GetZh() string { return c.Zh }
func (c nameableCity) GetEn() string { return c.En }
func (c nameableCounty) GetID() string { return c.ID }
func (c nameableCounty) GetZh() string { return c.Zh }
func (c nameableCounty) GetEn() string { return c.En }

func (d *Division) GetProvinceByID(id string) (Province, bool) {
    v, ok := d.provinceByID[id]
    return v, ok
}

func (d *Division) GetProvinceByName(name string) (Province, bool) {
    id, ok := d.provinceNameIndex[strings.ToLower(name)]
    if !ok {
        return Province{}, false
    }
    return d.provinceByID[id], true
}

func (d *Division) GetCityByID(id string) (City, bool) {
    v, ok := d.cityByID[id]
    return v, ok
}

func (d *Division) GetCityByName(name string) (City, bool) {
    id, ok := d.cityNameIndex[strings.ToLower(name)]
    if !ok {
        return City{}, false
    }
    return d.cityByID[id], true
}

func (d *Division) GetCountyByID(id string) (County, bool) {
    v, ok := d.countyByID[id]
    return v, ok
}

func (d *Division) GetCountyByName(name string) (County, bool) {
    id, ok := d.countyNameIndex[strings.ToLower(name)]
    if !ok {
        return County{}, false
    }
    return d.countyByID[id], true
}

func (d *Division) ListProvinces() []Province { return d.provinces }
func (d *Division) ListCities() []City { return d.cities }
func (d *Division) ListCounties() []County { return d.counties }

func (d *Division) GetCitiesByProvince(provinceID string) []City {
    r := make([]City, 0)
    prefix := provinceID[:2]
    for _, c := range d.cities {
        if hasPrefix(c.ID, prefix) {
            r = append(r, c)
        }
    }
    return r
}

func (d *Division) GetCountiesByCity(cityID string) []County {
    r := make([]County, 0)
    prefix := cityID[:4]
    for _, k := range d.counties {
        if hasPrefix(k.ID, prefix) {
            r = append(r, k)
        }
    }
    return r
}

type Item struct {
    Kind string
    ID string
    Zh string
    En string
}

func (d *Division) SearchByName(name string) []Item {
    q := strings.ToLower(name)
    r := make([]Item, 0)
    for _, p := range d.provinces {
        if strings.Contains(strings.ToLower(p.Zh), q) || strings.Contains(strings.ToLower(p.En), q) {
            r = append(r, Item{Kind: "province", ID: p.ID, Zh: p.Zh, En: p.En})
        }
    }
    for _, c := range d.cities {
        if strings.Contains(strings.ToLower(c.Zh), q) || strings.Contains(strings.ToLower(c.En), q) {
            r = append(r, Item{Kind: "city", ID: c.ID, Zh: c.Zh, En: c.En})
        }
    }
    for _, k := range d.counties {
        if strings.Contains(strings.ToLower(k.Zh), q) || strings.Contains(strings.ToLower(k.En), q) {
            r = append(r, Item{Kind: "county", ID: k.ID, Zh: k.Zh, En: k.En})
        }
    }
    return r
}
