package admtp

type Province struct {
    ID string `json:"id"`
    Zh string `json:"zh"`
    En string `json:"en"`
}

type City struct {
    ID string `json:"id"`
    Zh string `json:"zh"`
    En string `json:"en"`
}

type County struct {
    ID string `json:"id"`
    Zh string `json:"zh"`
    En string `json:"en"`
}

type provinceFile struct {
    Province []Province `json:"province"`
}

type citiesFile struct {
    Cities []City `json:"cities"`
}

type countiesFile struct {
    Counties []County `json:"counties"`
}
