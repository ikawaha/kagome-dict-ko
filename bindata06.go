package ko

import(
	"os"
	"time"
)


func dictKodictAgBytes() ([]byte, error) {
	return _dictKodictAg, nil
}

func dictKodictAg() (*asset, error) {
	bytes, err := dictKodictAgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.ag", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
