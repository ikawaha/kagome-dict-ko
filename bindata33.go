package ko

import(
	"os"
	"time"
)


func dictKodictBhBytes() ([]byte, error) {
	return _dictKodictBh, nil
}

func dictKodictBh() (*asset, error) {
	bytes, err := dictKodictBhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.bh", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
