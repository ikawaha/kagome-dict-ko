package ko

import(
	"os"
	"time"
)


func dictKodictAuBytes() ([]byte, error) {
	return _dictKodictAu, nil
}

func dictKodictAu() (*asset, error) {
	bytes, err := dictKodictAuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.au", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
