package ko

import(
	"os"
	"time"
)


func dictKodictAwBytes() ([]byte, error) {
	return _dictKodictAw, nil
}

func dictKodictAw() (*asset, error) {
	bytes, err := dictKodictAwBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.aw", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
