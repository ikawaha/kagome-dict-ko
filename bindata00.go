package ko

import(
	"os"
	"time"
)


func dictKodictAaBytes() ([]byte, error) {
	return _dictKodictAa, nil
}

func dictKodictAa() (*asset, error) {
	bytes, err := dictKodictAaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.aa", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
