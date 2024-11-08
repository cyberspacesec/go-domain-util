package domain_util

import (
	"github.com/weppos/publicsuffix-go/publicsuffix"
	"golang.org/x/net/idna"
	"strings"
	"unicode"
)

// SplitFldDomainAndSubDomain 将域名中的一级域名和其子域名分割开
func SplitFldDomainAndSubDomain(domain string) (fldDomain string, subDomainName string, err error) {
	domain = strings.ToLower(domain)
	punyCode := false
	if !isASCII(domain) {
		punyCode = true
	}
	if punyCode {
		domain, err = idna.ToASCII(domain)
	}
	if err != nil {
		return
	}
	d, err := publicsuffix.ParseFromListWithOptions(publicsuffix.DefaultList, domain,
		&publicsuffix.FindOptions{IgnorePrivate: true})
	if err != nil {
		return
	}

	if punyCode {
		d.TLD, err = idna.ToUnicode(d.TLD)
		if err != nil {
			return
		}
		d.SLD, err = idna.ToUnicode(d.SLD)
		if err != nil {
			return
		}
		d.TRD, err = idna.ToUnicode(d.TRD)
		if err != nil {
			return
		}
	}

	fldDomain = d.SLD + "." + d.TLD
	subDomainName = d.TRD
	return
}

// FldDomain 取域名中的一级域名部分
func FldDomain(domain string) (string, error) {
	fldDomain, _, err := SplitFldDomainAndSubDomain(domain)
	if err != nil {
		return "", err
	}
	return fldDomain, nil
}

// FldDomainIgnoreError 取域名中的一级域名部分，忽略错误，如果有错误的话则会返回空字符串
func FldDomainIgnoreError(domain string) string {
	fldDomain, _, _ := SplitFldDomainAndSubDomain(domain)
	return fldDomain
}

// IsFldDomain 判断给定的域名是否是一个根域名
func IsFldDomain(domain string) (bool, error) {
	_, subDomainName, err := SplitFldDomainAndSubDomain(domain)
	if err != nil {
		return false, err
	}
	return subDomainName == "", nil
}

// ---------------------------------------------------------------------------------------------------------------------

// 判断给定的字符串是否都会是ascii范围能够表示的
func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

// NewHostEntry 解析域名，分离其中的一级域名和子域名信息
func NewHostEntry(host string) (*HostEntry, error) {
	var err error
	host = strings.ToLower(host)
	punyCode := false
	if !IsASCII(host) {
		punyCode = true
	}
	origHost := host
	if punyCode {
		host, err = idna.ToASCII(host)
	}
	if err != nil {
		return nil, err
	}
	d, err := publicsuffix.ParseFromListWithOptions(publicsuffix.DefaultList, host,
		&publicsuffix.FindOptions{IgnorePrivate: true})
	if err != nil {
		return nil, err
	}

	if punyCode {
		d.TLD, err = idna.ToUnicode(d.TLD)
		if err != nil {
			return nil, err
		}
		d.SLD, err = idna.ToUnicode(d.SLD)
		if err != nil {
			return nil, err
		}
		d.TRD, err = idna.ToUnicode(d.TRD)
		if err != nil {
			return nil, err
		}
	}

	wBase := origHost[strings.Index(origHost, ".")+1:]
	if d.TRD == "" {
		wBase = d.SLD + "." + d.TLD
	}

	return &HostEntry{
		Host:         origHost,
		Tld:          d.TLD,
		Domain:       d.SLD + "." + d.TLD,
		SubName:      d.TRD,
		WildcardBase: wBase,
	}, nil
}

func IsASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

type HostEntry struct {
	Host            string
	Tld             string
	Domain          string
	SubName         string
	WildcardBase    string
	WildcardRecords []Record
	Wildcard        bool
}

type Record struct {
	Type  string `bson:"type"`
	Value string `bson:"value"`
}

// ------------------------------------------------- ------------------------------------------------------------------
