package ko

import(
	"os"
	"time"
)


func dictKodictBeBytes() ([]byte, error) {
	return _dictKodictBe, nil
}

func dictKodictBe() (*asset, error) {
	bytes, err := dictKodictBeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.be", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
