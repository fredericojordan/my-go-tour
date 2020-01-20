package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    pic := make([][]uint8, dy)

    for i := range pic {
        pic[i] = make([]uint8, dx)
    }

    for x := range pic {
        for y := range pic[x] {
            pic[x][y] = uint8((x+y)/2)
        }
    }

    return pic
}

func main() {
    pic.Show(Pic)
}
