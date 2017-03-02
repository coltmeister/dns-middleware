package vpndns

import (
    "os"
    "time"
    "io"
    "crypto/md5"
)


func watchFile(path string, delay int, action func(string) error) {
    previous := make([]byte, 0)

    for ; true; {
        current, err := hashFile(path)

        if err == nil {
            if !byteMatch(previous, current) {
                action(path)
            }
        }

        time.Sleep(time.Duration(delay) * time.Millisecond)
    }
}

func hashFile(path string) ([]byte, error) {
    h := md5.New()
    f, err := os.Open(path)

    if err != nil {
        return nil, err
    }

    defer f.Close()

    if _, err := io.Copy(h, f); err != nil {
        return nil, err
    }

    return h.Sum(nil), nil
}

func byteMatch(b1 []byte, b2 []byte) bool {
    if len(b1) != len(b2) {
        return false
    }

    for i, v := range b1 {
        if v != b2[i] {
            return false
        }
    }

    return true
}
