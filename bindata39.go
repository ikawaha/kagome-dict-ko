package ko

import(
	"os"
	"time"
)


func dictKodictBnBytes() ([]byte, error) {
	return _dictKodictBn, nil
}

func dictKodictBn() (*asset, error) {
	bytes, err := dictKodictBnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.bn", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
