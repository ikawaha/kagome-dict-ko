package ko

import(
	"os"
	"time"
)


func dictKodictAnBytes() ([]byte, error) {
	return _dictKodictAn, nil
}

func dictKodictAn() (*asset, error) {
	bytes, err := dictKodictAnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.an", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
