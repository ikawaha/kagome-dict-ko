package ko

import(
	"os"
	"time"
)


func dictKodictBdBytes() ([]byte, error) {
	return _dictKodictBd, nil
}

func dictKodictBd() (*asset, error) {
	bytes, err := dictKodictBdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.bd", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
