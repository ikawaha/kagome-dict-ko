package ko

import(
	"os"
	"time"
)


func dictKodictBsBytes() ([]byte, error) {
	return _dictKodictBs, nil
}

func dictKodictBs() (*asset, error) {
	bytes, err := dictKodictBsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.bs", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
