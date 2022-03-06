package html

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Valid_HTML_Tags(t *testing.T) {
	r := require.New(t)
	tags := []string{
		"a",
		"abbr",
		"address",
		"area",
		"article",
		"aside",
		"audio",
		"b",
		"base",
		"bdi",
		"bdo",
		"blockquote",
		"body",
		"br",
		"button",
		"canvas",
		"caption",
		"cite",
		"code",
		"col",
		"colgroup",
		"data",
		"datalist",
		"dd",
		"del",
		"details",
		"dfn",
		"dialog",
		"div",
		"dl",
		"dt",
		"em",
		"embed",
		"fieldset",
		"figcaption",
		"figure",
		"footer",
		"form",
		"h1",
		"head",
		"header",
		"hgroup",
		"hr",
		"html",
		"i",
		"iframe",
		"img",
		"input",
		"ins",
		"kbd",
		"label",
		"legend",
		"li",
		"link",
		"main",
		"map",
		"mark",
		"MathML",
		"menu",
		"meta",
		"meter",
		"nav",
		"noscript",
		"object",
		"ol",
		"optgroup",
		"option",
		"output",
		"p",
		"param",
		"picture",
		"pre",
		"progress",
		"q",
		"rp",
		"rt",
		"ruby",
		"s",
		"samp",
		"script",
		"section",
		"select",
		"slot",
		"small",
		"source",
		"span",
		"strong",
		"style",
		"sub",
		"summary",
		"sup",
		"svg",
		"table",
		"tbody",
		"td",
		"template",
		"textarea",
		"tfoot",
		"th",
		"thead",
		"time",
		"title",
		"tr",
		"track",
		"u",
		"ul",
		"var",
		"video",
		"wbr",
	}

	for _, tag := range tags {
		t.Run(fmt.Sprintf("Tag: %s", tag), func(t *testing.T) {
			got := IsValidHTMLTag(tag)
			r.True(got)
		})
	}
}

func Test_Invalid_HTML_Tags(t *testing.T) {
	r := require.New(t)
	tags := []string{
		"foo",
		"bar",
	}

	for _, tag := range tags {
		t.Run(fmt.Sprintf("Tag: %s", tag), func(t *testing.T) {
			got := IsValidHTMLTag(tag)
			r.False(got)
		})
	}
}

func Test_Valid_Global_HTML_Attributes(t *testing.T) {
	r := require.New(t)
	tags := []string{
		"a",
		"abbr",
		"address",
		"area",
		"article",
		"aside",
		"audio",
		"b",
		"base",
		"bdi",
		"bdo",
		"blockquote",
		"body",
		"br",
		"button",
		"canvas",
		"caption",
		"cite",
		"code",
		"col",
		"colgroup",
		"data",
		"datalist",
		"dd",
		"del",
		"details",
		"dfn",
		"dialog",
		"div",
		"dl",
		"dt",
		"em",
		"embed",
		"fieldset",
		"figcaption",
		"figure",
		"footer",
		"form",
		"h1",
		"head",
		"header",
		"hgroup",
		"hr",
		"html",
		"i",
		"iframe",
		"img",
		"input",
		"ins",
		"kbd",
		"label",
		"legend",
		"li",
		"link",
		"main",
		"map",
		"mark",
		"MathML",
		"menu",
		"meta",
		"meter",
		"nav",
		"noscript",
		"object",
		"ol",
		"optgroup",
		"option",
		"output",
		"p",
		"param",
		"picture",
		"pre",
		"progress",
		"q",
		"rp",
		"rt",
		"ruby",
		"s",
		"samp",
		"script",
		"section",
		"select",
		"slot",
		"small",
		"source",
		"span",
		"strong",
		"style",
		"sub",
		"summary",
		"sup",
		"svg",
		"table",
		"tbody",
		"td",
		"template",
		"textarea",
		"tfoot",
		"th",
		"thead",
		"time",
		"title",
		"tr",
		"track",
		"u",
		"ul",
		"var",
		"video",
		"wbr",
	}

	globalAttrs := []string{
		"accesskey",
		"autocapitalize",
		"autofocus",
		"contenteditable",
		"dir",
		"draggable",
		"enterkeyhint",
		"hidden",
		"inputmode",
		"is",
		"itemid",
		"itemprop",
		"itemref",
		"itemscope",
		"itemtype",
		"lang",
		"nonce",
		"spellcheck",
		"style",
		"tabindex",
		"title",
		"translate",
	}

	for _, tag := range tags {
		for _, attr := range globalAttrs {
			t.Run(fmt.Sprintf("Tag: %s, Attribute: %s", tag, attr), func(t *testing.T) {
				got, errCode := IsValidHTMLAttribute(tag, attr)
				r.True(got)
				r.Equal(ValidHtmlAttribute, errCode)
			})
		}
	}
}

func Test_Valid_HTML_Attributes(t *testing.T) {
	r := require.New(t)
	tests := []struct {
		tag       string
		attribute string
	}{
		{tag: "a", attribute: "href"},
		{tag: "a", attribute: "target"},
		{tag: "a", attribute: "download"},
		{tag: "a", attribute: "ping"},
		{tag: "a", attribute: "rel"},
		{tag: "a", attribute: "hreflang"},
		{tag: "a", attribute: "type"},
		{tag: "a", attribute: "referrerpolicy"},

		{tag: "area", attribute: "alt"},
		{tag: "area", attribute: "coords"},
		{tag: "area", attribute: "shape"},
		{tag: "area", attribute: "href"},
		{tag: "area", attribute: "target"},
		{tag: "area", attribute: "download"},
		{tag: "area", attribute: "ping"},
		{tag: "area", attribute: "rel"},
		{tag: "area", attribute: "referrerpolicy"},

		{tag: "audio", attribute: "src"},
		{tag: "audio", attribute: "crossorigin"},
		{tag: "audio", attribute: "preload"},
		{tag: "audio", attribute: "autoplay"},
		{tag: "audio", attribute: "loop"},
		{tag: "audio", attribute: "muted"},
		{tag: "audio", attribute: "controls"},

		{tag: "base", attribute: "href"},
		{tag: "base", attribute: "target"},

		{tag: "blockquote", attribute: "cite"},

		{tag: "body", attribute: "onafterprint"},
		{tag: "body", attribute: "onbeforeprint"},
		{tag: "body", attribute: "onbeforeunload"},
		{tag: "body", attribute: "onhashchange"},
		{tag: "body", attribute: "onlanguagechange"},
		{tag: "body", attribute: "onmessage"},
		{tag: "body", attribute: "onmessageerror"},
		{tag: "body", attribute: "onoffline"},
		{tag: "body", attribute: "ononline"},
		{tag: "body", attribute: "onpagehide"},
		{tag: "body", attribute: "onpageshow"},
		{tag: "body", attribute: "onpopstate"},
		{tag: "body", attribute: "onrejectionhandled"},
		{tag: "body", attribute: "onstorage"},
		{tag: "body", attribute: "onunhandledrejection"},
		{tag: "body", attribute: "onunload"},

		{tag: "button", attribute: "disabled"},
		{tag: "button", attribute: "form"},
		{tag: "button", attribute: "formaction"},
		{tag: "button", attribute: "formenctype"},
		{tag: "button", attribute: "formmethod"},
		{tag: "button", attribute: "formnovalidate"},
		{tag: "button", attribute: "formtarget"},
		{tag: "button", attribute: "name"},
		{tag: "button", attribute: "type"},
		{tag: "button", attribute: "value"},

		{tag: "canvas", attribute: "width"},
		{tag: "canvas", attribute: "height"},

		{tag: "col", attribute: "span"},

		{tag: "colgroup", attribute: "span"},

		{tag: "data", attribute: "value"},

		{tag: "del", attribute: "cite"},
		{tag: "del", attribute: "datetime"},

		{tag: "details", attribute: "open"},

		{tag: "dialog", attribute: "open"},

		{tag: "embed", attribute: "src"},
		{tag: "embed", attribute: "type"},
		{tag: "embed", attribute: "width"},
		{tag: "embed", attribute: "height"},

		{tag: "fieldset", attribute: "disabled"},
		{tag: "fieldset", attribute: "form"},
		{tag: "fieldset", attribute: "name"},

		{tag: "form", attribute: "accept-charset"},
		{tag: "form", attribute: "action"},
		{tag: "form", attribute: "autocomplete"},
		{tag: "form", attribute: "enctype"},
		{tag: "form", attribute: "method"},
		{tag: "form", attribute: "name"},
		{tag: "form", attribute: "novalidate"},
		{tag: "form", attribute: "target"},

		{tag: "html", attribute: "manifest"},

		{tag: "iframe", attribute: "src"},
		{tag: "iframe", attribute: "srcdoc"},
		{tag: "iframe", attribute: "name"},
		{tag: "iframe", attribute: "sandbox"},
		{tag: "iframe", attribute: "allow"},
		{tag: "iframe", attribute: "allowfullscreen"},
		{tag: "iframe", attribute: "width"},
		{tag: "iframe", attribute: "height"},
		{tag: "iframe", attribute: "referrerpolicy"},
		{tag: "iframe", attribute: "loading"},

		{tag: "img", attribute: "alt"},
		{tag: "img", attribute: "src"},
		{tag: "img", attribute: "srcset"},
		{tag: "img", attribute: "sizes"},
		{tag: "img", attribute: "crossorigin"},
		{tag: "img", attribute: "usemap"},
		{tag: "img", attribute: "ismap"},
		{tag: "img", attribute: "width"},
		{tag: "img", attribute: "height"},
		{tag: "img", attribute: "referrerpolicy"},
		{tag: "img", attribute: "decoding"},
		{tag: "img", attribute: "loading"},

		{tag: "input", attribute: "accept"},
		{tag: "input", attribute: "alt"},
		{tag: "input", attribute: "autocomplete"},
		{tag: "input", attribute: "checked"},
		{tag: "input", attribute: "dirname"},
		{tag: "input", attribute: "disabled"},
		{tag: "input", attribute: "form"},
		{tag: "input", attribute: "formaction"},
		{tag: "input", attribute: "formenctype"},
		{tag: "input", attribute: "formmethod"},
		{tag: "input", attribute: "formnovalidate"},
		{tag: "input", attribute: "formtarget"},
		{tag: "input", attribute: "height"},
		{tag: "input", attribute: "list"},
		{tag: "input", attribute: "max"},
		{tag: "input", attribute: "maxlength"},
		{tag: "input", attribute: "min"},
		{tag: "input", attribute: "minlength"},
		{tag: "input", attribute: "multiple"},
		{tag: "input", attribute: "name"},
		{tag: "input", attribute: "pattern"},
		{tag: "input", attribute: "placeholder"},
		{tag: "input", attribute: "readonly"},
		{tag: "input", attribute: "required"},
		{tag: "input", attribute: "size"},
		{tag: "input", attribute: "src"},
		{tag: "input", attribute: "step"},
		{tag: "input", attribute: "type"},
		{tag: "input", attribute: "value"},
		{tag: "input", attribute: "width"},

		{tag: "ins", attribute: "cite"},
		{tag: "ins", attribute: "datetime"},

		{tag: "label", attribute: "for"},

		{tag: "li", attribute: "value"},
		{tag: "li", attribute: "value"},

		{tag: "link", attribute: "href"},
		{tag: "link", attribute: "crossorigin"},
		{tag: "link", attribute: "rel"},
		{tag: "link", attribute: "as"},
		{tag: "link", attribute: "media"},
		{tag: "link", attribute: "hreflang"},
		{tag: "link", attribute: "type"},
		{tag: "link", attribute: "sizes"},
		{tag: "link", attribute: "imagesrcset"},
		{tag: "link", attribute: "imagesizes"},
		{tag: "link", attribute: "referrerpolicy"},
		{tag: "link", attribute: "integrity"},
		{tag: "link", attribute: "blocking"},
		{tag: "link", attribute: "color"},
		{tag: "link", attribute: "disabled"},

		{tag: "map", attribute: "name"},

		{tag: "meta", attribute: "name"},
		{tag: "meta", attribute: "http-equiv"},
		{tag: "meta", attribute: "content"},
		{tag: "meta", attribute: "charset"},
		{tag: "meta", attribute: "media"},

		{tag: "meter", attribute: "value"},
		{tag: "meter", attribute: "min"},
		{tag: "meter", attribute: "max"},
		{tag: "meter", attribute: "low"},
		{tag: "meter", attribute: "high"},
		{tag: "meter", attribute: "optimum"},

		{tag: "object", attribute: "data"},
		{tag: "object", attribute: "type"},
		{tag: "object", attribute: "name"},
		{tag: "object", attribute: "form"},
		{tag: "object", attribute: "width"},
		{tag: "object", attribute: "height"},

		{tag: "ol", attribute: "reversed"},
		{tag: "ol", attribute: "start"},
		{tag: "ol", attribute: "type"},

		{tag: "optgroup", attribute: "disabled"},
		{tag: "optgroup", attribute: "label"},

		{tag: "option", attribute: "disabled"},
		{tag: "option", attribute: "label"},
		{tag: "option", attribute: "selected"},
		{tag: "option", attribute: "value"},

		{tag: "output", attribute: "for"},
		{tag: "output", attribute: "form"},
		{tag: "output", attribute: "name"},

		{tag: "param", attribute: "name"},
		{tag: "param", attribute: "value"},

		{tag: "progress", attribute: "value"},
		{tag: "progress", attribute: "max"},

		{tag: "q", attribute: "cite"},

		{tag: "script", attribute: "src"},
		{tag: "script", attribute: "type"},
		{tag: "script", attribute: "async"},
		{tag: "script", attribute: "defer"},
		{tag: "script", attribute: "crossorigin"},
		{tag: "script", attribute: "integrity"},
		{tag: "script", attribute: "referrerpolicy"},
		{tag: "script", attribute: "blocking"},

		{tag: "select", attribute: "autocomplete"},
		{tag: "select", attribute: "disabled"},
		{tag: "select", attribute: "form"},
		{tag: "select", attribute: "multiple"},
		{tag: "select", attribute: "name"},
		{tag: "select", attribute: "required"},
		{tag: "select", attribute: "size"},

		{tag: "slot", attribute: "name"},

		{tag: "source", attribute: "src"},
		{tag: "source", attribute: "type"},
		{tag: "source", attribute: "srcset"},
		{tag: "source", attribute: "sizes"},
		{tag: "source", attribute: "media"},
		{tag: "source", attribute: "width"},
		{tag: "source", attribute: "height"},

		{tag: "style", attribute: "media"},
		{tag: "style", attribute: "blocking"},

		{tag: "td", attribute: "colspan"},
		{tag: "td", attribute: "rowspan"},
		{tag: "td", attribute: "headers"},

		{tag: "textarea", attribute: "cols"},
		{tag: "textarea", attribute: "dirname"},
		{tag: "textarea", attribute: "disabled"},
		{tag: "textarea", attribute: "form"},
		{tag: "textarea", attribute: "maxlength"},
		{tag: "textarea", attribute: "minlength"},
		{tag: "textarea", attribute: "name"},
		{tag: "textarea", attribute: "placeholder"},
		{tag: "textarea", attribute: "readonly"},
		{tag: "textarea", attribute: "required"},
		{tag: "textarea", attribute: "rows"},
		{tag: "textarea", attribute: "wrap"},

		{tag: "th", attribute: "colspan"},
		{tag: "th", attribute: "rowspan"},
		{tag: "th", attribute: "headers"},
		{tag: "th", attribute: "scope"},
		{tag: "th", attribute: "abbr"},

		{tag: "time", attribute: "datetime"},

		{tag: "track", attribute: "default"},
		{tag: "track", attribute: "kind"},
		{tag: "track", attribute: "label"},
		{tag: "track", attribute: "src"},
		{tag: "track", attribute: "srclang"},

		{tag: "video", attribute: "src"},
		{tag: "video", attribute: "crossorigin"},
		{tag: "video", attribute: "poster"},
		{tag: "video", attribute: "preload"},
		{tag: "video", attribute: "autoplay"},
		{tag: "video", attribute: "playsinline"},
		{tag: "video", attribute: "loop"},
		{tag: "video", attribute: "muted"},
		{tag: "video", attribute: "controls"},
		{tag: "video", attribute: "width"},
		{tag: "video", attribute: "height"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Tag: %s, Attribute: %s", tt.tag, tt.attribute), func(t *testing.T) {
			got, errCode := IsValidHTMLAttribute(tt.tag, tt.attribute)
			r.True(got)
			r.Equal(ValidHtmlAttribute, errCode)
		})
	}
}

func Test_Invalid_HTML_Attributes_ErrorInvalidHtmlTag(t *testing.T) {
	r := require.New(t)
	tests := []struct {
		tag       string
		attribute string
	}{
		{tag: "foo", attribute: "value"},
		{tag: "bar", attribute: "name"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Tag: %s, Attribute: %s", tt.tag, tt.attribute), func(t *testing.T) {
			got, errCode := IsValidHTMLAttribute(tt.tag, tt.attribute)
			r.False(got)
			r.Equal(ErrorInvalidHtmlTag, errCode)
		})
	}
}

func Test_Invalid_HTML_Attributes_ErrorInvalidOrUnsupportedHtmlAttribute(t *testing.T) {
	r := require.New(t)
	tests := []struct {
		tag       string
		attribute string
	}{
		{tag: "svg", attribute: "x"},
		{tag: "embed", attribute: "srcset"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Tag: %s, Attribute: %s", tt.tag, tt.attribute), func(t *testing.T) {
			got, errCode := IsValidHTMLAttribute(tt.tag, tt.attribute)
			r.False(got)
			r.Equal(ErrorInvalidOrUnsupportedHtmlAttribute, errCode)
		})
	}
}
