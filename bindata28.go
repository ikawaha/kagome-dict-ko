package ko

import(
	"os"
	"time"
)


func dictKodictBcBytes() ([]byte, error) {
	return _dictKodictBc, nil
}

func dictKodictBc() (*asset, error) {
	bytes, err := dictKodictBcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.bc", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
