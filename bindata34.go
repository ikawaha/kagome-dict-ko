package ko

import(
	"os"
	"time"
)


func dictKodictBiBytes() ([]byte, error) {
	return _dictKodictBi, nil
}

func dictKodictBi() (*asset, error) {
	bytes, err := dictKodictBiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.bi", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
