package ko

import(
	"os"
	"time"
)


func dictKodictBlBytes() ([]byte, error) {
	return _dictKodictBl, nil
}

func dictKodictBl() (*asset, error) {
	bytes, err := dictKodictBlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.bl", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
