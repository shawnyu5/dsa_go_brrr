package linkedlist_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLinkedlist(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Linkedlist Suite")
}
