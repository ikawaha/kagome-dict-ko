package ko

import(
	"os"
	"time"
)


func dictKodictAyBytes() ([]byte, error) {
	return _dictKodictAy, nil
}

func dictKodictAy() (*asset, error) {
	bytes, err := dictKodictAyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.ay", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
