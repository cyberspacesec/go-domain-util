package domain_util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFldDomain(t *testing.T) {

	domain, err := FldDomain("foo.bar.com")
	assert.Nil(t, err)
	assert.Equal(t, "bar.com", domain)

	domain, err = FldDomain("foo.bar.com.cn")
	assert.Nil(t, err)
	assert.Equal(t, "bar.com.cn", domain)

}

func TestFldDomainIgnoreError(t *testing.T) {
	domain := FldDomainIgnoreError("foo.bar.com")
	assert.Equal(t, "bar.com", domain)

	domain = FldDomainIgnoreError("foo.bar.com.cn")
	assert.Equal(t, "bar.com.cn", domain)
}

func TestIsFldDomain(t *testing.T) {

	ok, err := IsFldDomain("bar.com")
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = IsFldDomain("bar.com.cn")
	assert.Nil(t, err)
	assert.True(t, ok)

}

func TestSplitFldDomainAndSubDomain(t *testing.T) {

	fldDomain, subDomain, err := SplitFldDomainAndSubDomain("bar.com")
	assert.Nil(t, err)
	assert.Equal(t, "bar.com", fldDomain)
	assert.Equal(t, "", subDomain)

	fldDomain, subDomain, err = SplitFldDomainAndSubDomain("bar.com.cn")
	assert.Nil(t, err)
	assert.Equal(t, "bar.com.cn", fldDomain)
	assert.Equal(t, "", subDomain)

}
