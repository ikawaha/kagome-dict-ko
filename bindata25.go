package ko

import(
	"os"
	"time"
)


func dictKodictAzBytes() ([]byte, error) {
	return _dictKodictAz, nil
}

func dictKodictAz() (*asset, error) {
	bytes, err := dictKodictAzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.az", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
