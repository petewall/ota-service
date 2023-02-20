package test

import (
	"fmt"
	. "github.com/onsi/gomega/gbytes"
	"os/exec"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/petewall/ota-service/v2/lib"
	"github.com/phayes/freeport"
)

func TestTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Feature Test Suite")
}

var (
	client            lib.Client
	otaService        string
	otaServiceSession *gexec.Session
	otaServiceURL     string
)

var _ = BeforeSuite(func() {
	var err error
	otaService, err = gexec.Build("github.com/petewall/ota-service/v2")
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

var _ = BeforeEach(func() {
	port, err := freeport.GetFreePort()
	Expect(err).ToNot(HaveOccurred())
	otaServiceURL = fmt.Sprintf("http://localhost:%d", port)
	client = lib.Client{
		Host: otaServiceURL,
	}
	args := []string{
		"--port", fmt.Sprintf("%d", port),
	}
	command := exec.Command(otaService, args...)
	otaServiceSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())
	Eventually(otaServiceSession.Out, 10*time.Second).Should(Say("Listening on port"))

	Seed()
})

var _ = AfterEach(func() {
	otaServiceSession.Terminate().Wait()
	Eventually(otaServiceSession).Should(gexec.Exit())
})

func Seed() {
	//	Expect(client.AddFirmware("bootstrap", "1.0", []byte("bootstrap 1.0 firmware"))).To(Succeed())
	//	Expect(client.AddFirmware("bootstrap", "2.0", []byte("bootstrap 2.0 firmware"))).To(Succeed())
	//	Expect(client.AddFirmware("lightswitch", "2.0", []byte("lightswitch 2.0 firmware"))).To(Succeed())
}
