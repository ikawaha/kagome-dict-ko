package ko

import(
	"os"
	"time"
)


func dictKodictAiBytes() ([]byte, error) {
	return _dictKodictAi, nil
}

func dictKodictAi() (*asset, error) {
	bytes, err := dictKodictAiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/kodict.ai", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1595556976, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
