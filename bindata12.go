package ko

import(
	"os"
	"time"
)


func dictKodictAmBytes() ([]byte, error) {
	return _dictKodictAm, nil
}

func dictKodictAm() (*asset, error) {
	bytes, err := dictKodictAmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.am", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
