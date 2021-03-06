package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func main() {

	fmt.Println("Hello, playground")

	var pdf = gofpdf.New("P", "mm", "A4", "")

	pdf.SetXmpMetadata(XMP_HEADER)

	pdf.OutputFileAndClose("test.pdf")

}

var XMP_HEADER = []byte(`
        <?xpacket begin="ï»¿" id="W5M0MpCehiHzreSzNTczkc9d"?>
        <x:xmpmeta xmlns:x="adobe:ns:meta/">
        <rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#">
        <rdf:Description rdf:about="" xmlns:dc="http://purl.org/dc/elements/1.1/"><dc:title><rdf:Alt><rdf:li xml:lang="x-default" ></rdf:li></rdf:Alt></dc:title><dc:creator><rdf:Seq><rdf:li></rdf:li></rdf:Seq></dc:creator><dc:subject><rdf:Bag><rdf:li></rdf:li></rdf:Bag></dc:subject><dc:format>application/pdf</dc:format><dc:description><rdf:Alt><rdf:li xml:lang="x-default" ></rdf:li></rdf:Alt></dc:description></rdf:Description>
        <rdf:Description rdf:about="" xmlns:pdf="http://ns.adobe.com/pdf/1.3/"><pdf:Producer>iTextSharp 4.1.0 (based on iText 2.1.0)</pdf:Producer><pdf:Keywords></pdf:Keywords></rdf:Description>
        <rdf:Description rdf:about="" xmlns:xmp="http://ns.adobe.com/xap/1.0/"><xmp:ModifyDate>2020-03-13T08:44:31+01:00</xmp:ModifyDate><xmp:CreatorTool>Symtrax - Compleo Suite</xmp:CreatorTool><xmp:CreateDate>2020-03-13T08:44:31+01:00</xmp:CreateDate></rdf:Description>
        <rdf:Description rdf:about="" xmlns:pdfaid="http://www.aiim.org/pdfa/ns/id/"><pdfaid:part>3</pdfaid:part><pdfaid:conformance>A</pdfaid:conformance></rdf:Description>
        <rdf:Description xmlns:pdfaExtension="http://www.aiim.org/pdfa/ns/extension/" xmlns:pdfaSchema="http://www.aiim.org/pdfa/ns/schema#" xmlns:pdfaProperty="http://www.aiim.org/pdfa/ns/property#" rdf:about=""><pdfaExtension:schemas><rdf:Bag><rdf:li rdf:parseType="Resource"><pdfaSchema:schema>Factur-X PDFA Extension Schema</pdfaSchema:schema><pdfaSchema:namespaceURI>urn:factur-x:pdfa:CrossIndustryDocument:invoice:1p0#</pdfaSchema:namespaceURI><pdfaSchema:prefix>fx</pdfaSchema:prefix><pdfaSchema:property><rdf:Seq><rdf:li rdf:parseType="Resource"><pdfaProperty:name>DocumentFileName</pdfaProperty:name><pdfaProperty:valueType>Text</pdfaProperty:valueType><pdfaProperty:category>external</pdfaProperty:category><pdfaProperty:description>name of the embedded XML invoice file</pdfaProperty:description></rdf:li><rdf:li rdf:parseType="Resource"><pdfaProperty:name>DocumentType</pdfaProperty:name><pdfaProperty:valueType>Text</pdfaProperty:valueType><pdfaProperty:category>external</pdfaProperty:category><pdfaProperty:description>INVOICE</pdfaProperty:description></rdf:li><rdf:li rdf:parseType="Resource"><pdfaProperty:name>Version</pdfaProperty:name><pdfaProperty:valueType>Text</pdfaProperty:valueType> <pdfaProperty:category>external</pdfaProperty:category><pdfaProperty:description>The actual version of the Factur-X XML schema</pdfaProperty:description></rdf:li><rdf:li rdf:parseType="Resource"><pdfaProperty:name>ConformanceLevel</pdfaProperty:name><pdfaProperty:valueType>Text</pdfaProperty:valueType><pdfaProperty:category>external</pdfaProperty:category><pdfaProperty:description>The conformance level of the embedded Factur-X data</pdfaProperty:description></rdf:li></rdf:Seq></pdfaSchema:property></rdf:li></rdf:Bag></pdfaExtension:schemas></rdf:Description><rdf:Description xmlns:fx="urn:factur-x:pdfa:CrossIndustryDocument:invoice:1p0#" rdf:about="" fx:ConformanceLevel="EN 16931" fx:DocumentFileName="factur-x.xml" fx:DocumentType="INVOICE" fx:Version="1.0"/>
    </rdf:RDF></x:xmpmeta>
    <?xpacket end="w"?>`)
