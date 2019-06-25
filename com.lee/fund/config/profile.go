package config

import (
	"encoding/xml"
	"fmt"
	"os"
)

type profileConfig struct {
	active   []string
	settings SettingMap
}

func (c *profileConfig) filter(s string) string {
	if c.settings == nil {
		return s
	}

	return os.Expand(s, func(n string) string {
		if v, ok := c.settings[n]; ok {
			return v
		}
		return "${" + n + "}"
	})
}

func (c *profileConfig) load(content []byte) error {
	type Profiles struct {
		Profiles [] struct {
			Name     string `xml:"name,attr"`
			Settings [] struct {
				Key   string `xml:"key,attr"`
				Value string `xml:"value,attr"`
			} `xml:"setting"`
		} `xml:"profile"`
	}

	var profiles Profiles
	err := xml.Unmarshal(content, &profiles)
	if err != nil {
		return fmt.Errorf("parse profile.conf failed：%v", err)
	}

	fn := func(pa string) {
		for _, pf := range profiles.Profiles {
			if pf.Name == pa {
				for _, s := range pf.Settings {
					c.settings[s.Key] = s.Value
				}
				break
			}
		}
	}

	c.settings = SettingMap{}
	if c.active == nil || len(c.active) == 0 {
		fn("")
	} else {
		for _, n := range c.active {
			fn(n)
		}
	}

	if len(c.settings) > 0 {
		fmt.Println("config > load profile settings...")
		for k, v := range c.settings {
			fmt.Printf("\t%v=%v\n", k, v)
		}
	} else {
		fmt.Println("config > load profile settings fail...")
	}

	return nil
}
