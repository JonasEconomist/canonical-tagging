package canonical

type root struct {
	cogito *cogito `xml:" cogito,omitempty" json:"cogito,omitempty"`
}

type cogito struct {
	doc *doc `xml:" doc,omitempty" json:"doc,omitempty"`
}

type doc struct {
	content   *content     `xml:" content,omitempty" json:"content,omitempty"`
	knowledge []*knowledge `xml:" knowledge,omitempty" json:"knowledge,omitempty"`
}

type content struct {
	text *text `xml:" text,omitempty" json:"text,omitempty"`
}

type text struct {
	Attr_charset  string `xml:" charset,attr"  json:",omitempty"`
	Attr_mimetype string `xml:" mimetype,attr"  json:",omitempty"`
	Text          string `xml:",chardata" json:",omitempty"`
}

type knowledge struct {
	Attr_name   string       `xml:" name,attr"  json:",omitempty"`
	annotations *annotations `xml:" annotations,omitempty" json:"annotations,omitempty"`
	descriptors *descriptors `xml:" descriptors,omitempty" json:"descriptors,omitempty"`
	types       *types       `xml:" types,omitempty" json:"types,omitempty"`
}

type types struct {
	tagType []*tagType `xml:" type,omitempty" json:"type,omitempty"`
}

type tagType struct {
	Attr_fullname string     `xml:" fullname,attr"  json:",omitempty"`
	Attr_name     string     `xml:" name,attr"  json:",omitempty"`
	ad            []*ad      `xml:" ad,omitempty" json:"ad,omitempty"`
	tagType       []*tagType `xml:" type,omitempty" json:"type,omitempty"`
}

type ad struct {
	Attr_name  string `xml:" name,attr"  json:",omitempty"`
	Attr_scope string `xml:" scope,attr"  json:",omitempty"`
	Attr_type  string `xml:" type,attr"  json:",omitempty"`
}

type descriptors struct {
	descriptor []*descriptor `xml:" descriptor,omitempty" json:"descriptor,omitempty"`
}

type descriptor struct {
	Attr_label string `xml:" label,attr"  json:",omitempty"`
	Attr_name  string `xml:" name,attr"  json:",omitempty"`
	Attr_type  string `xml:" type,attr"  json:",omitempty"`
	a          []*a   `xml:" a,omitempty" json:"a,omitempty"`
}

type a struct {
	Attr_name  string `xml:" name,attr"  json:",omitempty"`
	Attr_type  string `xml:" type,attr"  json:",omitempty"`
	Attr_value string `xml:" value,attr"  json:",omitempty"`
}

type annotations struct {
	annotation []*annotation `xml:" annotation,omitempty" json:"annotation,omitempty"`
}

type annotation struct {
	Attr_e    string `xml:" e,attr"  json:",omitempty"`
	Attr_name string `xml:" name,attr"  json:",omitempty"`
	Attr_s    string `xml:" s,attr"  json:",omitempty"`
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	a         []*a   `xml:" a,omitempty" json:"a,omitempty"`
}
