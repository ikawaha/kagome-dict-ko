package ko

import(
	"os"
	"time"
)


func dictKodictAtBytes() ([]byte, error) {
	return _dictKodictAt, nil
}

func dictKodictAt() (*asset, error) {
	bytes, err := dictKodictAtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.at", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
