package ko

import(
	"os"
	"time"
)


func dictKodictBaBytes() ([]byte, error) {
	return _dictKodictBa, nil
}

func dictKodictBa() (*asset, error) {
	bytes, err := dictKodictBaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.ba", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
