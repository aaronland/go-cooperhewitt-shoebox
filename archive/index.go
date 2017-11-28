package archive

type ArchiveIndex struct {
     Html HTML `xml:"html"`
}

type ArchiveIndexHTML struct {
     Head ArchiveIndexHead   `xml:"head"`
     Body ArchiveIndexBoxy   `xml:"body"`
}

type ArchiveIndexHead struct {
     Title string `xml:"title"`
     Meta  []ArchiveIndexMeta `xml:"meta"`
     Link  []ArchiveIndexLink `xml:"link"`     
     Script  []ArchiveIndexScript `xml:"script"`     
}

type ArchiveIndexMeta struct {
     Name string `xml:"name"`
     Content string `xml:"content"`
     HTTPEquiv string `xml:"http-equiv"`
}

type ArchiveIndexLink struct {
     Rel string `xml:"rel"`
     Type string `xml:"type"`
     Href string `xml:"href"`
}

type ArchiveIndexScript struct {
     Src string `xml:"src"`
     Type string `xml:"type"`
     // body? 
}

type ArchiveIndexBody struct {
}
