package admtp

import (
    "encoding/json"
    "errors"
    "os"
    "path/filepath"
    "runtime"
)

func getDataDirectory() (string, error) {
    _, thisFile, _, _ := runtime.Caller(0)
    base := filepath.Dir(thisFile)
    p := filepath.Join(base, "data")
    if st, err := os.Stat(p); err == nil && st.IsDir() {
        return p, nil
    }
    return "", errors.New("data directory not found")
}

func loadJSON[T any](path string, v *T) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()
    dec := json.NewDecoder(f)
    return dec.Decode(v)
}

func loadProvinces() ([]Province, error) {
    dir, err := getDataDirectory()
    if err != nil {
        return nil, err
    }
    var pf provinceFile
    if err := loadJSON(filepath.Join(dir, "province.json"), &pf); err != nil {
        return nil, err
    }
    return pf.Province, nil
}

func loadCities(provinceID string) ([]City, error) {
    dir, err := getDataDirectory()
    if err != nil {
        return nil, err
    }
    var cf citiesFile
    if err := loadJSON(filepath.Join(dir, "cities", provinceID+".json"), &cf); err != nil {
        return nil, err
    }
    return cf.Cities, nil
}

func loadCounties(provinceID string) ([]County, error) {
    dir, err := getDataDirectory()
    if err != nil {
        return nil, err
    }
    var kf countiesFile
    if err := loadJSON(filepath.Join(dir, "counties", provinceID, "counties.json"), &kf); err != nil {
        return nil, err
    }
    return kf.Counties, nil
}
