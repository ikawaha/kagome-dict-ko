package ko

import(
	"os"
	"time"
)


func dictKodictAhBytes() ([]byte, error) {
	return _dictKodictAh, nil
}

func dictKodictAh() (*asset, error) {
	bytes, err := dictKodictAhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.ah", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
