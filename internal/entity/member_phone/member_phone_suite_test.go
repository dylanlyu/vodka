package member_phone_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMemberPhone(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MemberPhone Suite")
}
