package ko

import(
	"os"
	"time"
)


func dictKodictBgBytes() ([]byte, error) {
	return _dictKodictBg, nil
}

func dictKodictBg() (*asset, error) {
	bytes, err := dictKodictBgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.bg", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
