package ko

import(
	"os"
	"time"
)


func dictKodictBbBytes() ([]byte, error) {
	return _dictKodictBb, nil
}

func dictKodictBb() (*asset, error) {
	bytes, err := dictKodictBbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.bb", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
