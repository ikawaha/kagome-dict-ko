package ko

import(
	"os"
	"time"
)


func dictKodictAoBytes() ([]byte, error) {
	return _dictKodictAo, nil
}

func dictKodictAo() (*asset, error) {
	bytes, err := dictKodictAoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.ao", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
