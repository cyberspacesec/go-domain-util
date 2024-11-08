package domain_util

//import "strings"
//
//// EmailDomain 解析邮箱所关联到的域名，指的是将@后面的部分取出来
//// 比如输入： foo@a.bar.com，会返回a.bar.com
//func EmailDomain(email string) string {
//	index := strings.LastIndex(email, "@")
//	if index < 0 {
//		return ""
//	}
//	return email[index+1:]
//}
//
//// EmailFldDomain 解析邮箱所关联到的一级域名
//// 比如输入： foo@a.bar.com，会返回bar.com，而不是a.bar.com
//func EmailFldDomain(email string) string {
//	domain := EmailDomain(email)
//	if domain == "" {
//		return ""
//	}
//	entry, err := NewHostEntry(domain)
//	if err != nil {
//		return ""
//	}
//	return entry.Domain
//}
