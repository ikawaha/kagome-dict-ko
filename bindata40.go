package ko

import(
	"os"
	"time"
)


func dictKodictBoBytes() ([]byte, error) {
	return _dictKodictBo, nil
}

func dictKodictBo() (*asset, error) {
	bytes, err := dictKodictBoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.bo", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
