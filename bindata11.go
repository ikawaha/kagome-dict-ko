package ko

import(
	"os"
	"time"
)


func dictKodictAlBytes() ([]byte, error) {
	return _dictKodictAl, nil
}

func dictKodictAl() (*asset, error) {
	bytes, err := dictKodictAlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.al", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
