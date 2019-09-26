package cat

import (
	"fmt"
	"github.com/yishuixing/xslt-cat/common"
	"io/ioutil"
	"regexp"
	"strings"
)

const (
	pattern_import      = `<xsl:import href="(\w+\.xslt)`
	pattern_stylesheet  = `<xsl:stylesheet[^>]+>|\<\?xml[^>]+>|<xsl:import href="(\w+.xslt)"\s*/>`
	pattern_stylesheet2 = `<xsl:stylesheet[^>]+>`
)

func Cat(xslt_dir string) {
	dm_content := common.ReadToString(xslt_dir + `\dm.xslt`)
	items := GetXslt(dm_content, pattern_import)
	dm_c := CleanXslt(dm_content)
	all_xslt := `<?xml version="1.0" encoding="utf-8"?>
<xsl:stylesheet version="1.0" 
                xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
                xmlns:msxsl="urn:schemas-microsoft-com:xslt"                 
                xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
                xmlns:dc="http://www.purl.org/dc/elements/1.1/" 
                  xmlns:xlink="http://www.w3.org/1999/xlink"
                exclude-result-prefixes="msxsl">`
	all_xslt += dm_c
	for _, f := range items {
		c := common.ReadToString(xslt_dir + `\` + f)
		all_xslt += CleanXslt(c)
	}
	all_xslt += `</xsl:stylesheet>`
	ioutil.WriteFile("DM_S1000D_4.2.xslt", []byte(all_xslt), 0644)
	fmt.Println("已生成 DM_S1000D_4.2.xslt")
}
func GetXslt(s, pattern string) []string {
	reg := regexp.MustCompile(pattern)
	submatch := reg.FindAllSubmatch([]byte(s), -1)
	var items []string
	for _, v := range submatch {
		for i, v1 := range v {
			s1 := string(v1)
			if i == 1 {
				items = append(items, s1)
			}
		}
	}
	return items
}
func CleanXslt(s string) string {
	reg := regexp.MustCompile(pattern_stylesheet)
	result := reg.ReplaceAllString(s, "")
	result = strings.ReplaceAll(result, `</xsl:stylesheet>`, "")
	return result
}
