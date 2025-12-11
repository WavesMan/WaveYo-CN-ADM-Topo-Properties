package admtp

import "strings"

func buildProvinceNameIndex(items []Province) map[string]string {
    idx := make(map[string]string)
    for _, it := range items {
        idx[strings.ToLower(it.Zh)] = it.ID
        idx[strings.ToLower(it.En)] = it.ID
        switch it.Zh {
        case "北京":
            idx[strings.ToLower("北京市")] = it.ID
        case "天津":
            idx[strings.ToLower("天津市")] = it.ID
        case "上海":
            idx[strings.ToLower("上海市")] = it.ID
        case "重庆":
            idx[strings.ToLower("重庆市")] = it.ID
        case "香港":
            idx[strings.ToLower("香港特别行政区")] = it.ID
        case "澳门":
            idx[strings.ToLower("澳门特别行政区")] = it.ID
        case "西藏":
            idx[strings.ToLower("西藏自治区")] = it.ID
        case "内蒙古":
            idx[strings.ToLower("内蒙古自治区")] = it.ID
        default:
            if !strings.HasSuffix(it.Zh, "省") && !strings.HasSuffix(it.Zh, "市") && !strings.HasSuffix(it.Zh, "自治区") && !strings.HasSuffix(it.Zh, "特别行政区") {
                idx[strings.ToLower(it.Zh+"省")] = it.ID
            }
        }
    }
    return idx
}

func buildNameIndex[T interface{ GetID() string; GetZh() string; GetEn() string }](items []T) map[string]string {
    idx := make(map[string]string)
    for _, it := range items {
        idx[strings.ToLower(it.GetZh())] = it.GetID()
        idx[strings.ToLower(it.GetEn())] = it.GetID()
    }
    return idx
}

func hasPrefix(s, prefix string) bool {
    return strings.HasPrefix(s, prefix)
}
