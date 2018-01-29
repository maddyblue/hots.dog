package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/mjibson/hots.dog/casc/xmlpath"
)

func (x *XML) loadXML(name string) error {
	if x.v == nil {
		x.v = make(map[string]*xmlpath.Node)
		x.ids = make(map[string]map[string]bool)
		x.parents = make(map[string]map[string]bool)
	}
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}
	dec := xml.NewDecoder(bytes.NewReader(b))
	dec.CharsetReader = func(charset string, r io.Reader) (io.Reader, error) {
		return r, nil
	}
	v, err := xmlpath.ParseDecoder(dec)
	if err != nil {
		return err
	}
	x.v[name] = v
	// Get all nodes with an id attr.
	iter := idNode.Iter(v)
	for iter.Next() {
		node := iter.Node()
		// Fetch its id.
		res, _ := nodeID.String(node)
		if len(x.ids[res]) == 0 {
			x.ids[res] = make(map[string]bool)
		}
		x.ids[res][name] = true
		// See if it has a parent.
		parent, ok := nodeParent.String(node)
		if ok {
			if len(x.parents[res]) == 0 {
				x.parents[res] = make(map[string]bool)
			}
			x.parents[res][parent] = true
		}
	}

	return nil
}

type XML struct {
	v map[string]*xmlpath.Node
	// ids holds a map from ID to nodes
	ids map[string]map[string]bool
	// parents holds a map from child ID to parent ID
	parents map[string]map[string]bool
}

var (
	reNum       = regexp.MustCompile(`\[.+?\]`)
	idNode      = xmlpath.MustCompile("//*[@id]")
	nodeID      = xmlpath.MustCompile("@id")
	nodeParent  = xmlpath.MustCompile("@parent")
	pathDefault = xmlpath.MustCompile("@default")
)

func (x *XML) Get(s string) (string, error) {
	def := "0"
	// I have no idea about this. See Button/Tooltip/AnaAimDownSightsCustomOptics.
	if s == "Value" {
		return def, nil
	}
	sp := strings.Split(s, ",")
	if len(sp) != 3 {
		return "", errors.Errorf("got %d splits: %q", len(sp), s)
	}
	typ := sp[0]
	id := sp[1]
	name := sp[2]
	origName := sp[2]
	name = reNum.ReplaceAllStringFunc(name, func(r string) string {
		r = r[1 : len(r)-1]
		i, err := strconv.Atoi(r)
		if err != nil {
			return fmt.Sprintf("[@index='%s']", r)
		}
		i++
		return fmt.Sprintf("[%d]", i)
	})
	name = strings.TrimRight(name, ".")
	name = strings.Replace(name, ".", "/", -1)
	sp = strings.Split(name, "/")
	first := strings.Join(sp[:len(sp)-1], "/")
	if len(first) > 0 {
		first += "/"
	}
	last := sp[len(sp)-1]
	var names []string
	for n := range x.ids[id] {
		names = append(names, n)
	}
	idpath := xmlpath.MustCompile(fmt.Sprintf("//*[@id='%s']", id))
	paths := []string{
		fmt.Sprintf("%s/@value", name),
		fmt.Sprintf("%s@%s", first, last),
	}
	var xps []*xmlpath.Path
	for _, p := range paths {
		xps = append(xps, xmlpath.MustCompile(p))
	}

	for _, name := range names {
		iditer := idpath.Iter(x.v[name])
		for iditer.Next() {
			idnode := iditer.Node()
			idname := idnode.Name()
			if !strings.HasPrefix(idname, "C"+typ) {
				continue
			}
			for _, xp := range xps {
				iter := xp.Iter(idnode)
				for iter.Next() {
					res := iter.Node().String()
					return res, nil
				}
			}
			iter := pathDefault.Iter(idnode)
			for iter.Next() {
				def = iter.Node().String()
			}
		}
	}

	for parent := range x.parents[id] {
		res, err := x.Get(fmt.Sprintf("%s,%s,%s", typ, parent, origName))
		if err != nil {
			return "", err
		}
		if res != "0" {
			return res, nil
		}
	}

	return def, nil
}
