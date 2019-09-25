package mode

import (
	"github.com/pkg/errors"
)

var (
	ModeSetError = errors.New("mode setting error, ['pro','dev','local','sta','test']")
)

type ModeType string

var (
	ModePro   = ModeType("pro")
	ModeDev   = ModeType("dev")
	ModeLocal = ModeType("local")
	ModeSta   = ModeType("sta")
	ModeTest  = ModeType("test")

	ModeList        = []ModeType{ModePro, ModeDev, ModeLocal, ModeSta, ModeTest}
	ModesOnlineList = []ModeType{ModePro, ModeDev, ModeSta, ModeTest}
	modeS           string
	modeT           ModeType
)

func InModesList(md ModeType, li []ModeType) bool {
	for _, m := range li {
		if md == m {
			return true
		}
	}
	return false
}

func InModesOnline(md ModeType) bool {
	return InModesList(md, ModesOnlineList)

}

func InModes(md ModeType) bool {
	return InModesList(md, ModeList)
}

func SetMode(m string) error {
	modeS, modeT = m, ModeType(m)
	if !InModes(modeT) {
		return ModeSetError
	}
	return nil
}

func GetModeString() string {
	return modeS
}

func GetMode() ModeType {
	return modeT
}
