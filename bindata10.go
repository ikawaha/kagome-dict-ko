package ko

import(
	"os"
	"time"
)


func dictKodictAkBytes() ([]byte, error) {
	return _dictKodictAk, nil
}

func dictKodictAk() (*asset, error) {
	bytes, err := dictKodictAkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.ak", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
