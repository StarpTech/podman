package e2e_test

import (
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("podman machine set", func() {
	var (
		mb      *machineTestBuilder
		testDir string
	)

	BeforeEach(func() {
		testDir, mb = setup()
	})
	AfterEach(func() {
		teardown(originalHomeDir, testDir, mb)
	})

	It("set machine cpus, disk, memory", func() {
		name := randomString()
		i := new(initMachine)
		session, err := mb.setName(name).setCmd(i.withImagePath(mb.imagePath)).run()
		Expect(err).To(BeNil())
		Expect(session).To(Exit(0))

		set := setMachine{}
		setSession, err := mb.setName(name).setCmd(set.withCPUs(2).withDiskSize(102).withMemory(4096)).run()
		Expect(err).To(BeNil())
		Expect(setSession).To(Exit(0))

		// shrinking disk size is verboten
		shrink, err := mb.setName(name).setCmd(set.withDiskSize(5)).run()
		Expect(err).To(BeNil())
		Expect(shrink).To(Exit(125))

		s := new(startMachine)
		startSession, err := mb.setCmd(s).run()
		Expect(err).To(BeNil())
		Expect(startSession).To(Exit(0))

		sshCPU := sshMachine{}
		CPUsession, err := mb.setName(name).setCmd(sshCPU.withSSHComand([]string{"lscpu", "|", "grep", "\"CPU(s):\"", "|", "head", "-1"})).run()
		Expect(err).To(BeNil())
		Expect(CPUsession).To(Exit(0))
		Expect(CPUsession.outputToString()).To(ContainSubstring("2"))

		sshDisk := sshMachine{}
		diskSession, err := mb.setName(name).setCmd(sshDisk.withSSHComand([]string{"sudo", "fdisk", "-l", "|", "grep", "Disk"})).run()
		Expect(err).To(BeNil())
		Expect(diskSession).To(Exit(0))
		Expect(diskSession.outputToString()).To(ContainSubstring("102 GiB"))

		sshMemory := sshMachine{}
		memorySession, err := mb.setName(name).setCmd(sshMemory.withSSHComand([]string{"cat", "/proc/meminfo", "|", "grep", "-i", "'memtotal'", "|", "grep", "-o", "'[[:digit:]]*'"})).run()
		Expect(err).To(BeNil())
		Expect(memorySession).To(Exit(0))
		foundMemory, err := strconv.Atoi(memorySession.outputToString())
		Expect(err).To(BeNil())
		Expect(foundMemory).To(BeNumerically(">", 3800000))
		Expect(foundMemory).To(BeNumerically("<", 4200000))

		// Setting a running machine results in 125
		runner, err := mb.setName(name).setCmd(set.withCPUs(4)).run()
		Expect(err).To(BeNil())
		Expect(runner).To(Exit(125))
	})

	It("no settings should change if no flags", func() {
		name := randomString()
		i := new(initMachine)
		session, err := mb.setName(name).setCmd(i.withImagePath(mb.imagePath)).run()
		Expect(err).To(BeNil())
		Expect(session).To(Exit(0))

		set := setMachine{}
		setSession, err := mb.setName(name).setCmd(&set).run()
		Expect(err).To(BeNil())
		Expect(setSession).To(Exit(0))

		s := new(startMachine)
		startSession, err := mb.setCmd(s).run()
		Expect(err).To(BeNil())
		Expect(startSession).To(Exit(0))

		ssh2 := sshMachine{}
		sshSession2, err := mb.setName(name).setCmd(ssh2.withSSHComand([]string{"lscpu", "|", "grep", "\"CPU(s):\"", "|", "head", "-1"})).run()
		Expect(err).To(BeNil())
		Expect(sshSession2).To(Exit(0))
		Expect(sshSession2.outputToString()).To(ContainSubstring("1"))

		ssh3 := sshMachine{}
		sshSession3, err := mb.setName(name).setCmd(ssh3.withSSHComand([]string{"sudo", "fdisk", "-l", "|", "grep", "Disk"})).run()
		Expect(err).To(BeNil())
		Expect(sshSession3).To(Exit(0))
		Expect(sshSession3.outputToString()).To(ContainSubstring("100 GiB"))
	})

})
