package ko

import(
	"os"
	"time"
)


func dictKodictBjBytes() ([]byte, error) {
	return _dictKodictBj, nil
}

func dictKodictBj() (*asset, error) {
	bytes, err := dictKodictBjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.bj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
