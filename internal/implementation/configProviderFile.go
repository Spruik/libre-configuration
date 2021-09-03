package implementation

import (
	"errors"
	"fmt"
	"github.com/Spruik/libre-configuration/shared"
	"github.com/antchfx/jsonquery"
	"os"
)

type configProviderFile struct {
	doc      *jsonquery.Node
	topicMap map[string]*jsonquery.Node
}

func NewConfigProviderFile() *configProviderFile {
	return &configProviderFile{
		topicMap: map[string]*jsonquery.Node{},
	}
}

func (s *configProviderFile) Initialize(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("Cannot find config file at %s err= %+v", filePath, err))
	}
	s.doc, err = jsonquery.Parse(f)
	if err != nil {
		panic(fmt.Sprintf("Cannot parse config file at %s  err= %+v", filePath, err))
	}

	nodes, err := jsonquery.QueryAll(s.doc, "/*")
	for _, j := range nodes {
		s.topicMap[j.Data] = j
	}
}

func (s *configProviderFile) GetComponentConfig(component string) *shared.ConfigItem {
	node := s.topicMap[component]
	if node != nil {
		ret := s.processNode(node)
		return ret
	} else {
		return &shared.ConfigItem{}
	}

}

func (s *configProviderFile) GetConfigEntry(component string, key string) (string, error) {
	comp := s.topicMap[component]
	if comp != nil {
		node, err := jsonquery.Query(comp, key)
		if err == nil {
			if node != nil {
				return node.InnerText(), nil
			} else {
				return "", errors.New(fmt.Sprintf("Key given is not in the component configuration: %s/%s", component, key))
			}
		} else {
			return "", err
		}
	} else {
		return "", errors.New(fmt.Sprintf("Component name given is not in the configuration: %s", component))
	}
}

func (s *configProviderFile) GetConfigEntryWithDefault(component string, key string, dflt string) (string, error) {
	comp := s.topicMap[component]
	if comp != nil {
		node, err := jsonquery.Query(comp, key)
		if err == nil {
			if node != nil {
				return node.InnerText(), nil
			} else {
				return "", nil
			}
		} else {
			return dflt, nil
		}
	} else {
		return "", errors.New(fmt.Sprintf("Component name given is not in the configuration: %s", component))
	}
}

func (s *configProviderFile) GetConfigStanza(component string, top string) (*shared.ConfigItem, error) {
	node, err := jsonquery.Query(s.topicMap[component], top)
	if err == nil {
		if node != nil {
			ret := s.processNode(node)
			return ret, nil
		} else {
			return &shared.ConfigItem{}, nil
		}
	}
	return nil, err
}

func (s *configProviderFile) processNode(currNode *jsonquery.Node) *shared.ConfigItem {
	ret := shared.ConfigItem{}
	ret.Name = currNode.Data
	for child := currNode.FirstChild; child != nil; child = child.NextSibling {
		switch child.Type {
		case jsonquery.ElementNode:
			ret.Children = append(ret.Children, s.processNode(child))
		case jsonquery.TextNode:
			ret.Value = child.Data
		}
	}
	return &ret
}
