package coordinate

import (
    "math"
)

const (
    xPi = 3.14159265358979324 * 3000.0 / 180.0
    // pi
    pi = 3.1415926535897932384626
    // 长半轴
    a = 6378245.0
    // 扁率
    ee = 0.00669342162296594323
)

func transformLatitude(longitude, latitude float64) float64 {
    ret := -100.0 + 2.0 * longitude + 3.0 * latitude + 0.2 * latitude * latitude +  0.1 * longitude * latitude + 0.2 * math.Sqrt(math.Abs(longitude))
    ret += (20.0 * math.Sin(6.0 * longitude * pi) + 20.0 * math.Sin(2.0 * longitude * pi)) * 2.0 / 3.0
    ret += (20.0 * math.Sin(latitude * pi) + 40.0 * math.Sin(latitude / 3.0 * pi)) * 2.0 / 3.0
    ret += (160.0 * math.Sin(latitude / 12.0 * pi) + 320 * math.Sin(latitude * pi / 30.0)) * 2.0 / 3.0
    return ret
}

func transformLongitude(longitude, latitude float64) float64 {
    ret := 300.0 + longitude + 2.0 * latitude + 0.1 * longitude * longitude + 0.1 * longitude * latitude + 0.1 * math.Sqrt(math.Abs(longitude))
    ret += (20.0 * math.Sin(6.0 * longitude * pi) + 20.0 * math.Sin(2.0 * longitude * pi)) * 2.0 / 3.0
    ret += (20.0 * math.Sin(longitude * pi) + 40.0 * math.Sin(longitude / 3.0 * pi)) * 2.0 / 3.0
    ret += (150.0 * math.Sin(longitude / 12.0 * pi) + 300.0 * math.Sin(longitude / 30.0 * pi)) * 2.0 / 3.0
    return ret
}

func isOversea(longitude, latitude float64) bool {
    if longitude < 72.004 || longitude > 137.8347 {
        return true
    }
    if latitude < 0.8293 || latitude > 55.8271 {
        return true
    }
    return false
}

func Wgs84ToGcj02(longitude, latitude float64) (float64, float64) {
    if isOversea(longitude, latitude) {
        return longitude, latitude
    }

    dLat := transformLatitude(longitude - 105.0, latitude - 35.0)
    dLng := transformLongitude(longitude - 105.0, latitude - 35.0)
    radLat := latitude / 180.0 * pi
    magic := math.Sin(radLat)
    magic = 1 - ee * magic * magic
    sqrtMagic := math.Sqrt(magic)
    dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * pi)
    dLng = (dLng * 180.0) / (a / sqrtMagic * math.Cos(radLat) * pi)
    return longitude + dLng, latitude + dLat
}

func Wgs84ToBd09(longitude, latitude float64) (float64, float64) {
    gcjLng, gcjLat := Wgs84ToGcj02(longitude, latitude)
    return Gcj02ToBd09(gcjLng, gcjLat)
}

func Gcj02ToBd09(longitude, latitude float64) (float64, float64) {
    z := math.Sqrt(longitude * longitude + latitude * latitude) + 0.00002 * math.Sin(latitude * xPi)
    theta := math.Atan2(latitude, longitude) + 0.000003 * math.Cos(longitude * xPi)
    return z * math.Cos(theta) + 0.0065, z * math.Sin(theta) + 0.006
}

func Gcj02ToWgs84(longitude, latitude float64) (float64, float64) {
    if isOversea(longitude, latitude) {
        return longitude, latitude
    }

    dLat := transformLatitude(longitude - 105.0, latitude - 35.0)
    dLng := transformLongitude(longitude - 105.0, latitude - 35.0)
    radLat := latitude / 180.0 * pi
    magic := math.Sin(radLat)
    magic = 1 - ee * magic * magic
    sqrtMagic := math.Sqrt(magic)
    dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * pi)
    dLng = (dLng * 180.0) / (a / sqrtMagic * math.Cos(radLat) * pi)
    mgLat := latitude + dLat
    mgLng := longitude + dLng
    return longitude * 2 - mgLng, latitude * 2 - mgLat
}

func Bd09ToGcj02(longitude, latitude float64) (float64, float64) {
    x := longitude - 0.0065
    y := latitude - 0.006
    z := math.Sqrt(x * x + y * y) - 0.00002 * math.Sin(y * xPi)
    theta := math.Atan2(y, x) - 0.000003 * math.Cos(x * xPi)
    return z * math.Cos(theta), z * math.Sin(theta)
}


func Bd09ToWgs84(longitude, latitude float64) (float64, float64) {
    gcjLng, gcjLat := Bd09ToGcj02(longitude, latitude)
    return Gcj02ToWgs84(gcjLng, gcjLat)
}
