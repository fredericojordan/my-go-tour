package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
    n, err = r.r.Read(b)

    for i := 0; i<n; i++ {
        if b[i] >= 65 && b[i] <= 122 {
            b[i] += 13

            if ((b[i] > 90 && b[i] < 110) || b[i] > 122) {
                b[i] -= 26
            }
        }
    }

    return
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
