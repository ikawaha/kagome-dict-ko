package ko

import(
	"os"
	"time"
)


func dictKodictAxBytes() ([]byte, error) {
	return _dictKodictAx, nil
}

func dictKodictAx() (*asset, error) {
	bytes, err := dictKodictAxBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.ax", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
