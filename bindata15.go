package ko

import(
	"os"
	"time"
)


func dictKodictApBytes() ([]byte, error) {
	return _dictKodictAp, nil
}

func dictKodictAp() (*asset, error) {
	bytes, err := dictKodictApBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.ap", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
