package ko

import(
	"os"
	"time"
)


func dictKodictBkBytes() ([]byte, error) {
	return _dictKodictBk, nil
}

func dictKodictBk() (*asset, error) {
	bytes, err := dictKodictBkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.bk", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
