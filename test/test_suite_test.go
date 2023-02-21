package test_test

import (
	"fmt"
	. "github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/ghttp"
	"os/exec"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/petewall/ota-service/lib"
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

	deviceService   *ghttp.Server
	firmwareService *ghttp.Server
)

var _ = BeforeSuite(func() {
	var err error
	otaService, err = gexec.Build("github.com/petewall/ota-service")
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

var _ = BeforeEach(func() {
	deviceService = ghttp.NewServer()
	firmwareService = ghttp.NewServer()

	port, err := freeport.GetFreePort()
	Expect(err).ToNot(HaveOccurred())
	args := []string{
		"--port", fmt.Sprintf("%d", port),
		"--device-service", deviceService.URL(),
		"--firmware-service", firmwareService.URL(),
	}
	command := exec.Command(otaService, args...)
	otaServiceSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())
	Eventually(otaServiceSession.Out, 10*time.Second).Should(Say("Listening on port"))

	client = lib.Client{
		Host: fmt.Sprintf("http://localhost:%d", port),
	}
})

var _ = AfterEach(func() {
	otaServiceSession.Terminate().Wait()
	Eventually(otaServiceSession).Should(gexec.Exit())

	deviceService.Close()
	firmwareService.Close()
})

func Seed() {
	//	Expect(client.AddFirmware("bootstrap", "1.0", []byte("bootstrap 1.0 firmware"))).To(Succeed())
	//	Expect(client.AddFirmware("bootstrap", "2.0", []byte("bootstrap 2.0 firmware"))).To(Succeed())
	//	Expect(client.AddFirmware("lightswitch", "2.0", []byte("lightswitch 2.0 firmware"))).To(Succeed())
}
