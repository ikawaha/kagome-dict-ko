package ko

import(
	"os"
	"time"
)


func dictKodictAqBytes() ([]byte, error) {
	return _dictKodictAq, nil
}

func dictKodictAq() (*asset, error) {
	bytes, err := dictKodictAqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.aq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
