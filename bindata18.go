package ko

import(
	"os"
	"time"
)


func dictKodictAsBytes() ([]byte, error) {
	return _dictKodictAs, nil
}

func dictKodictAs() (*asset, error) {
	bytes, err := dictKodictAsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.as", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
