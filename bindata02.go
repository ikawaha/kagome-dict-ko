package ko

import(
	"os"
	"time"
)


func dictKodictAcBytes() ([]byte, error) {
	return _dictKodictAc, nil
}

func dictKodictAc() (*asset, error) {
	bytes, err := dictKodictAcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.ac", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
