package implementation

import (
	"errors"
	"fmt"
	"os"

	"github.com/Spruik/libre-configuration/shared"
	"github.com/antchfx/jsonquery"
)

const errMessageNotInitalized = "configProviderFile not initialized. Please call `Initialize(\"path/to/file.json\")`"

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
		panic(fmt.Sprintf("cannot find config file at %s err= %+v", filePath, err))
	}
	s.doc, err = jsonquery.Parse(f)
	if err != nil {
		panic(fmt.Sprintf("cannot parse config file at %s  err= %+v", filePath, err))
	}

	nodes, err := jsonquery.QueryAll(s.doc, "/*")
	if err != nil {
		panic(fmt.Errorf("failed to json deserialize file %s, expected no error; got %s", filePath, err))
	}
	for _, j := range nodes {
		s.topicMap[j.Data] = j
	}
}

func (s *configProviderFile) GetComponentConfig(component string) *shared.ConfigItem {
	if s == nil {
		return nil
	}
	node := s.topicMap[component]
	if node != nil {
		ret := s.processNode(node)
		return ret
	} else {
		return &shared.ConfigItem{}
	}

}

func (s *configProviderFile) GetConfigEntry(component string, key string) (string, error) {
	if s == nil {
		return "", errors.New(errMessageNotInitalized)
	}
	comp := s.topicMap[component]
	if comp != nil {
		node, err := jsonquery.Query(comp, key)
		if err == nil {
			if node != nil {
				return node.InnerText(), nil
			} else {
				return "", fmt.Errorf("key given is not in the component configuration: %s/%s", component, key)
			}
		} else {
			return "", err
		}
	} else {
		return "", fmt.Errorf("component name given is not in the configuration: %s", component)
	}
}

func (s *configProviderFile) GetConfigEntryWithDefault(component string, key string, dflt string) (string, error) {
	if s == nil {
		return dflt, errors.New(errMessageNotInitalized)
	}
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
		return "", fmt.Errorf("component name given is not in the configuration: %s", component)
	}
}

func (s *configProviderFile) GetConfigStanza(component string, top string) (*shared.ConfigItem, error) {
	if s == nil {
		return nil, errors.New(errMessageNotInitalized)
	}
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
