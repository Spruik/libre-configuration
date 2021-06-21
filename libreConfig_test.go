package libreConfig

import (
	"github.com/Spruik/libre-configuration/shared"
	"log"
	"testing"
)

func TestInitialize(t *testing.T) {
	Initialize("./TestConfig.json")
}

func TestGetConfigEntry(t *testing.T) {
	cat := "category3"
	key := "stanza1/stanzastanza1/stanzaitem2"
	val, err := GetConfigService().GetConfigEntry(cat, key)
	if err == nil {
		log.Printf("%s/%s=%s", cat, key, val)
	} else {
		t.Logf("%s", err)
		t.Fail()
	}
}

func TestGetConfigStanza(t *testing.T) {
	val, err := GetConfigService().GetConfigStanza("libreLogger", "loggers")
	if err == nil {
		log.Print("libreLogger/loggers")
		printItem(val, "")
	} else {
		t.Logf("%s", err)
		t.Fail()
	}
}

func TestGetConfigStanzaForList(t *testing.T) {
	val, err := GetConfigService().GetConfigStanza("configWithList", "MyList")
	if err == nil {
		log.Print("configWithList/MyList")
		printItem(val, "")
	} else {
		t.Logf("%s", err)
		t.Fail()
	}
}

func printItem(item *shared.ConfigItem, indent string) {
	log.Printf("%sName: \"%s\"  Value: \"%s\"", indent, item.Name, item.Value)
	for _, child := range item.Children {
		printItem(child, indent+"   ")
	}
}
