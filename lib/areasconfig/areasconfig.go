package areasconfig

import (
	"errors"
	"github.com/askovpen/goated/lib/config"
	"github.com/askovpen/goated/lib/msgapi"
	//	"log"
	"sort"
	"strings"
)

// Read area configs
func Read() error {
	switch config.Config.AreaFile.Type {
	case "fidoconfig":
		err := fidoConfigRead(config.Config.AreaFile.Path)
		if err != nil {
			return err
		}
	default:
		return errors.New("unknown AreasConfig.Type '" + config.Config.AreaFile.Type + "'")
	}
	for i := range config.Config.Areas {
		found := false
		for _, da := range msgapi.Areas {
			if config.Config.Areas[i].Name == da.GetName() {
				found = true
				if config.Config.Areas[i].Chrs != "" {
					da.SetChrs(config.Config.Areas[i].Chrs)
				}
			}
		}
		if !found {
			a, err := getArea(i)
			if err == nil {
				msgapi.Areas = append(msgapi.Areas, a)
			}
		}
	}

	sort.Slice(msgapi.Areas, func(i, j int) bool {
		if msgapi.Areas[i].GetType() != msgapi.Areas[j].GetType() {
			return msgapi.Areas[i].GetType() < msgapi.Areas[j].GetType()
		}
		return msgapi.Areas[i].GetName() < msgapi.Areas[j].GetName()
	})
	//	log.Printf("%q",msgapi.Areas)
	return nil
}

func getArea(i int) (msgapi.AreaPrimitive, error) {
	switch config.Config.Areas[i].BaseType {
	case "msg":
		r := &msgapi.MSG{AreaName: config.Config.Areas[i].Name, AreaPath: config.Config.Areas[i].Path, AreaType: getType(config.Config.Areas[i].Type)}
		if config.Config.Areas[i].Chrs != "" {
			r.Chrs = config.Config.Areas[i].Chrs
		}
		return r, nil
	case "squish":
		r := &msgapi.Squish{AreaName: config.Config.Areas[i].Name, AreaPath: config.Config.Areas[i].Path, AreaType: getType(config.Config.Areas[i].Type)}
		if config.Config.Areas[i].Chrs != "" {
			r.Chrs = config.Config.Areas[i].Chrs
		}
		return r, nil
	case "jam":
		r := &msgapi.JAM{AreaName: config.Config.Areas[i].Name, AreaPath: config.Config.Areas[i].Path, AreaType: getType(config.Config.Areas[i].Type)}
		if config.Config.Areas[i].Chrs != "" {
			r.Chrs = config.Config.Areas[i].Chrs
		}
		return r, nil
	}
	return nil, errors.New("uknown type")
}
func getType(t string) msgapi.EchoAreaType {
	if strings.EqualFold(t, "echo") {
		return msgapi.EchoAreaTypeEcho
	} else if strings.EqualFold(t, "local") {
		return msgapi.EchoAreaTypeLocal
	} else if strings.EqualFold(t, "netmail") {
		return msgapi.EchoAreaTypeNetmail
	}
	return msgapi.EchoAreaTypeLocal
}
