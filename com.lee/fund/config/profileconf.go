package config

import (
	"encoding/xml"
	"fmt"
	"os"
)

type profileConf struct {
	active   []string
	settings SettingMap
}

type profile struct {
	Name     string `xml:"name,attr"`
	Settings [] struct {
		Key   string `xml:"key,attr"`
		Value string `xml:"value,attr"`
	} `xml:"setting"`
}

func (pcf *profileConf) filter(s string) string {
	if pcf.settings == nil {
		return s
	}

	return os.Expand(s, func(n string) string {
		if v, ok := pcf.settings[n]; ok {
			return v
		}
		return "${" + n + "}"
	})
}

func (pcf *profileConf) load(content []byte) error {
	pf := struct {
		Profiles [] *profile `xml:"profile"`
	}{}

	err := xml.Unmarshal(content, &pf)
	if err != nil {
		return fmt.Errorf("parse profile.conf failed：%v", err)
	}

	fn := func(pa string) {
		for _, pf := range pf.Profiles {
			if pf.Name == pa {
				for _, s := range pf.Settings {
					pcf.settings[s.Key] = s.Value
				}
				break
			}
		}
	}

	pcf.settings = SettingMap{}
	if pcf.active == nil || len(pcf.active) == 0 {
		fn("")
	} else {
		for _, n := range pcf.active {
			fn(n)
		}
	}

	if len(pcf.settings) > 0 {
		fmt.Println("config > load profile settings...")
		for k, v := range pcf.settings {
			fmt.Printf("\t%v=%v\n", k, v)
		}
	} else {
		fmt.Println("config > load profile settings fail...")
	}

	return nil
}
