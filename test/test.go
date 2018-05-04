package main

import (
    "fmt"
    "github.com/hwsdien/coordinate"
)

func main() {
    wgsLng, wgsLat := coordinate.Bd09ToWgs84(113.273881382906, 23.1290303717461)
    fmt.Printf("%g,%g\n", wgsLat, wgsLng)

    gcjLng, gcjLat:= coordinate.Bd09ToGcj02(113.273881382906, 23.1290303717461)
    fmt.Printf("%g,%g\n", gcjLat, gcjLng)


    a := 22.011834
    b := 113.370867

    bdLng, bdLat := coordinate.Wgs84ToBd09(b, a)
    fmt.Printf("%g,%g\n", bdLat, bdLng)
}
