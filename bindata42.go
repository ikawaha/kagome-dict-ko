package ko

import(
	"os"
	"time"
)


func dictKodictBqBytes() ([]byte, error) {
	return _dictKodictBq, nil
}

func dictKodictBq() (*asset, error) {
	bytes, err := dictKodictBqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.bq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
