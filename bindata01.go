package ko

import(
	"os"
	"time"
)


func dictKodictAbBytes() ([]byte, error) {
	return _dictKodictAb, nil
}

func dictKodictAb() (*asset, error) {
	bytes, err := dictKodictAbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.ab", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
