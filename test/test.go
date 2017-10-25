package main

import (
    "fmt"
    "github.com/hwsdien/coordinate"
)

func main() {
    wgsLng, wgsLat := coordinate.Bd09ToWgs84(116.404, 39.915)
    fmt.Printf("%g,%g\n", wgsLat, wgsLng)
    bdLng, bdLat := coordinate.Wgs84ToBd09(wgsLng, wgsLat)
    fmt.Printf("%g,%g\n", bdLat, bdLng)
}
