package ko

import(
	"os"
	"time"
)


func dictKodictAvBytes() ([]byte, error) {
	return _dictKodictAv, nil
}

func dictKodictAv() (*asset, error) {
	bytes, err := dictKodictAvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.av", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
