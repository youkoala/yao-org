package core

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/net/html"
)

// BindEvent is a method that binds events to the page.
func (page *Page) BindEvent(ctx *BuildContext, sel *goquery.Selection, component string, ispage bool) {
	matcher := NewAttrPrefixMatcher(`s:on-`)
	sel.FindMatcher(matcher).Each(func(i int, s *goquery.Selection) {
		if comp, has := s.Attr("is"); has && ctx.isJitComponent(comp) {
			return
		}
		script := GetEventScript(&ctx.sequence, s, page.namespace, component, ispage)
		if script != nil {
			ctx.scripts = append(ctx.scripts, *script)
		}
	})
}

// GetEventScript the event script
func GetEventScript(sequence *int, sel *goquery.Selection, ns string, cn string, ispage bool) *ScriptNode {

	if len(sel.Nodes) == 0 {
		return nil
	}

	// Page events
	events := map[string]string{}
	dataUnique := map[string]string{}
	jsonUnique := map[string]string{}
	id := fmt.Sprintf("event-%d", *sequence)
	*sequence = *sequence + 1

	for _, attr := range sel.Nodes[0].Attr {

		if strings.HasPrefix(attr.Key, "s:on-") {
			name := strings.TrimPrefix(attr.Key, "s:on-")
			handler := attr.Val
			events[name] = handler
			continue
		}

		if strings.HasPrefix(attr.Key, "s:data-") {
			name := strings.TrimPrefix(attr.Key, "s:data-")
			dataUnique[name] = attr.Val
			sel.SetAttr(fmt.Sprintf("data:%s", name), attr.Val)
			continue
		}

		if strings.HasPrefix(attr.Key, "s:json-") {
			name := strings.TrimPrefix(attr.Key, "s:json-")
			jsonUnique[name] = attr.Val
			sel.SetAttr(fmt.Sprintf("json:%s", name), attr.Val)
			continue
		}
	}

	data := []string{}
	for name := range dataUnique {
		data = append(data, name)
		sel.RemoveAttr(fmt.Sprintf("s:data-%s", name))
	}

	json := []string{}
	for name := range jsonUnique {
		json = append(json, name)
		sel.RemoveAttr(fmt.Sprintf("s:json-%s", name))
	}

	dataRaw, _ := jsoniter.MarshalToString(data)
	jsonRaw, _ := jsoniter.MarshalToString(json)

	source := ""
	for name, handler := range events {
		if ispage {
			source += pageEventInjectScript(id, name, dataRaw, jsonRaw, handler) + "\n"
		} else {
			source += compEventInjectScript(id, name, cn, dataRaw, jsonRaw, handler) + "\n"
		}
		sel.RemoveAttr(fmt.Sprintf("s:on-%s", name))
	}

	sel.SetAttr("s:event", id)

	return &ScriptNode{
		Source:    source,
		Namespace: ns,
		Component: cn,
		Attrs:     []html.Attribute{{Key: "event", Val: id}},
	}
}
