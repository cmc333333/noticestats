package stats

import (
	"gopkg.in/xmlpath.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func fetchXML(model *NoticeJSON) string {
	resp, err := http.Get(model.FullTextXMLUrl)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}
	return string(body)
}

func regtextLen(xml string) int {
	length := 0
	path := xmlpath.MustCompile("//REGTEXT")
	root, err := xmlpath.Parse(strings.NewReader(xml))
	if err != nil {
		log.Panic(err)
	}
	nodes := path.Iter(root)
	for nodes.Next() {
		length += len(nodes.Node().String())
	}
	return length
}
