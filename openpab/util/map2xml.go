package util

import (
	"encoding/xml"
	"io"
)

type FileRoot map[string]string

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (m FileRoot) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}

	return e.EncodeToken(start.End())
}

func (m *FileRoot) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = FileRoot{}
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}

func Map2Xml(m map[string]string) (string, error) {
	buf, err := xml.Marshal(FileRoot(m))
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func Xml2Map(str string) (map[string]string, error) {
	m := make(map[string]string)
	if err := xml.Unmarshal([]byte(str), (*FileRoot)(&m)); err != nil {
		return nil, nil
	}

	return m, nil
}
