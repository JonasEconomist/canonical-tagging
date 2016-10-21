

package main

/////////////////////////////////////////////////////////////////
//Code generated by chidley https://github.com/gnewton/chidley //
/////////////////////////////////////////////////////////////////

import (
	"bufio"
	"compress/bzip2"
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
)

const (
	JsonOut = iota
	XmlOut
	CountAll
)

var toJson bool = false
var toXml bool = false
var oneLevelDown bool = false
var countAll bool = false
var musage bool = false

var uniqueFlags = []*bool{
	&toJson,
	&toXml,
	&countAll}

var filename = "/Users/kathrynjonas/Work/src/github.com/gnewton/chidley/xml/fullGo.xml"



func init() {
	flag.BoolVar(&toJson, "j", toJson, "Convert to JSON")
	flag.BoolVar(&toXml, "x", toXml, "Convert to XML")
	flag.BoolVar(&countAll, "c", countAll, "Count each instance of XML tags")
	flag.BoolVar(&oneLevelDown, "s", oneLevelDown, "Stream XML by using XML elements one down from the root tag. Good for huge XML files (see http://blog.davidsingleton.org/parsing-huge-xml-files-with-go/")
	flag.BoolVar(&musage, "h", musage, "Usage")
	flag.StringVar(&filename, "f", filename, "XML file or URL to read in")
}

var out int = -1

var counters map[string]*int

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	if musage {
		flag.Usage()
		return
	}

	numSetBools, outFlag := numberOfBoolsSet(uniqueFlags)
	if numSetBools == 0 {
		flag.Usage()
                return
	}

	if numSetBools != 1 {
		flag.Usage()
		log.Fatal("Only one of ", uniqueFlags, " can be set at once")
	}

	reader, xmlFile, err := genericReader(filename)
	if err != nil {
		log.Fatal(err)
		return
	}

	decoder := xml.NewDecoder(reader)
	counters = make(map[string]*int)
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		switch se := token.(type) {
		case xml.StartElement:
			handleFeed(se, decoder, outFlag)
		}
	}
        if xmlFile != nil{
	    defer xmlFile.Close()
        }
	if countAll {
		for k, v := range counters {
			fmt.Println(*v, k)
		}
	}
}

func handleFeed(se xml.StartElement, decoder *xml.Decoder, outFlag *bool) {
	if outFlag == &countAll {
		incrementCounter(se.Name.Space, se.Name.Local)
	} else {
                if !oneLevelDown{
        		if se.Name.Local == "cogito" && se.Name.Space == "" {
	        	      var item Chicogito
			      decoder.DecodeElement(&item, &se)
			      switch outFlag {
			      case &toJson:
				      writeJson(item)
			      case &toXml:
				      writeXml(item)
			      }
		      }
                }else{
                   
        		if se.Name.Local == "doc" && se.Name.Space == "" {
	        	      var item Chidoc
			      decoder.DecodeElement(&item, &se)
			      switch outFlag {
			      case &toJson:
				      writeJson(item)
			      case &toXml:
				      writeXml(item)
			      }
		      }
                   
               }
	}
}

func makeKey(space string, local string) string {
	if space == "" {
		space = "_"
	}
	return space + ":" + local
}

func incrementCounter(space string, local string) {
	key := makeKey(space, local)

	counter, ok := counters[key]
	if !ok {
		n := 1
		counters[key] = &n
	} else {
		newv := *counter + 1
		counters[key] = &newv
	}
}

func writeJson(item interface{}) {
	b, err := json.MarshalIndent(item, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

func writeXml(item interface{}) {
	output, err := xml.MarshalIndent(item, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write(output)
}

func genericReader(filename string) (io.Reader, *os.File, error) {
	if filename == "" {
		return bufio.NewReader(os.Stdin), nil, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	if strings.HasSuffix(filename, "bz2") {
		return bufio.NewReader(bzip2.NewReader(bufio.NewReader(file))), file, err
	}

	if strings.HasSuffix(filename, "gz") {
		reader, err := gzip.NewReader(bufio.NewReader(file))
		if err != nil {
			return nil, nil, err
		}
		return bufio.NewReader(reader), file, err
	}
	return bufio.NewReader(file), file, err
}

func numberOfBoolsSet(a []*bool) (int, *bool) {
	var setBool *bool
	counter := 0
	for i := 0; i < len(a); i++ {
		if *a[i] {
			counter += 1
			setBool = a[i]
		}
	}
	return counter, setBool
}


///////////////////////////
/// structs
///////////////////////////

type Chiroot struct {
	Chicogito *Chicogito `xml:" cogito,omitempty" json:"cogito,omitempty"`
}

type Chicogito struct {
	Chidoc *Chidoc `xml:" doc,omitempty" json:"doc,omitempty"`
}

type Chidoc struct {
	Chicontent *Chicontent `xml:" content,omitempty" json:"content,omitempty"`
	Chiknowledge []*Chiknowledge `xml:" knowledge,omitempty" json:"knowledge,omitempty"`
}

type Chicontent struct {
	Chitext *Chitext `xml:" text,omitempty" json:"text,omitempty"`
}

type Chitext struct {
	Attr_charset string `xml:" charset,attr"  json:",omitempty"`
	Attr_mimetype string `xml:" mimetype,attr"  json:",omitempty"`
	Text string `xml:",chardata" json:",omitempty"`
}

type Chiknowledge struct {
	Attr_name string `xml:" name,attr"  json:",omitempty"`
	Chiannotations *Chiannotations `xml:" annotations,omitempty" json:"annotations,omitempty"`
	Chidescriptors *Chidescriptors `xml:" descriptors,omitempty" json:"descriptors,omitempty"`
	Chitypes *Chitypes `xml:" types,omitempty" json:"types,omitempty"`
}

type Chitypes struct {
	Chitype []*Chitype `xml:" type,omitempty" json:"type,omitempty"`
}

type Chitype struct {
	Attr_fullname string `xml:" fullname,attr"  json:",omitempty"`
	Attr_name string `xml:" name,attr"  json:",omitempty"`
	Chiad []*Chiad `xml:" ad,omitempty" json:"ad,omitempty"`
	Chitype []*Chitype `xml:" type,omitempty" json:"type,omitempty"`
}

type Chiad struct {
	Attr_name string `xml:" name,attr"  json:",omitempty"`
	Attr_scope string `xml:" scope,attr"  json:",omitempty"`
	Attr_type string `xml:" type,attr"  json:",omitempty"`
}

type Chidescriptors struct {
	Chidescriptor []*Chidescriptor `xml:" descriptor,omitempty" json:"descriptor,omitempty"`
}

type Chidescriptor struct {
	Attr_label string `xml:" label,attr"  json:",omitempty"`
	Attr_name string `xml:" name,attr"  json:",omitempty"`
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Chia []*Chia `xml:" a,omitempty" json:"a,omitempty"`
}

type Chia struct {
	Attr_name string `xml:" name,attr"  json:",omitempty"`
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Attr_value string `xml:" value,attr"  json:",omitempty"`
}

type Chiannotations struct {
	Chiannotation []*Chiannotation `xml:" annotation,omitempty" json:"annotation,omitempty"`
}

type Chiannotation struct {
	Attr_e string `xml:" e,attr"  json:",omitempty"`
	Attr_name string `xml:" name,attr"  json:",omitempty"`
	Attr_s string `xml:" s,attr"  json:",omitempty"`
	Attr_type string `xml:" type,attr"  json:",omitempty"`
	Chia []*Chia `xml:" a,omitempty" json:"a,omitempty"`
}


///////////////////////////
