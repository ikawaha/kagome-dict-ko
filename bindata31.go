package ko

import(
	"os"
	"time"
)


func dictKodictBfBytes() ([]byte, error) {
	return _dictKodictBf, nil
}

func dictKodictBf() (*asset, error) {
	bytes, err := dictKodictBfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.bf", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
