package ko

import(
	"os"
	"time"
)


func dictKodictAfBytes() ([]byte, error) {
	return _dictKodictAf, nil
}

func dictKodictAf() (*asset, error) {
	bytes, err := dictKodictAfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.af", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
