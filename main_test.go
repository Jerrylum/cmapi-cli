package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CommandSpec struct {
	ExpectedCommandLine string
	MockOutput          string
	MockError           string
	MockExitCode        int
}

var MockCommandsQueue = []CommandSpec{}

func mockExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperMockExecute", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}

	if len(MockCommandsQueue) > 0 {
		item := MockCommandsQueue[0]
		cmd.Env = append(cmd.Env, "EXPECTED_COMMAND_LINE="+item.ExpectedCommandLine)
		cmd.Env = append(cmd.Env, "MOCK_OUTPUT="+item.MockOutput)
		cmd.Env = append(cmd.Env, "MOCK_ERROR="+item.MockError)
		cmd.Env = append(cmd.Env, "MOCK_EXIT_CODE="+fmt.Sprintf("%d", item.MockExitCode))
		MockCommandsQueue = MockCommandsQueue[1:]
	} else {
		panic("MockCommandsQueue is empty")
	}
	return cmd
}

func TestHelperMockExecute(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	expectedCommandLine := os.Getenv("EXPECTED_COMMAND_LINE")
	wantedOutput := os.Getenv("MOCK_OUTPUT")
	wantedError := os.Getenv("MOCK_ERROR")
	wantedExitCode, err := strconv.Atoi(os.Getenv("MOCK_EXIT_CODE"))

	commandLine := strings.Join(os.Args[3:], " ")

	if commandLine != expectedCommandLine || err != nil {
		os.Exit(127)
		return
	}

	fmt.Fprintf(os.Stdout, "%s", wantedOutput)
	fmt.Fprintf(os.Stderr, "%s", wantedError)
	os.Exit(wantedExitCode)
}

func setup() {
	ExecCommand = mockExecCommand
}

func teardown() {
	os.Remove("data.json")
	os.Remove("project.pros")
	os.Remove(".cmapi-cli-secret.json")
	ExecCommand = exec.Command
	MockCommandsQueue = []CommandSpec{}
}

func TestReadWriteJson(t *testing.T) {
	setup()
	defer teardown()

	data := make(map[string]string)
	data["key1"] = ""
	data["key2"] = "value2"
	data["key3"] = "123"
	data[""] = "value4"

	WriteJson("data.json", data)
	data2 := ReadJson("data.json")

	assert.Equal(t, data, data2)
}

func TestSetupSecret(t *testing.T) {
	setup()
	defer teardown()

	wd, _ := os.Getwd()

	AdminDir = wd
	assert.True(t, SetupSecret())
	assert.NotEqual(t, "unknown", Secret["computer-name"])

	file, _ := os.Stat(SecretFilePath)
	assert.NotNil(t, file)

	os.Remove(".cmapi-cli-secret.json")

	file, _ = os.Stat(SecretFilePath)
	assert.Nil(t, file)
}

func TestUpdateFileSecret(t *testing.T) {
	setup()
	defer teardown()

	setA := make(map[string]string)
	setA["key1"] = "value1"
	setA["key2"] = "value2"
	setA["key3"] = "value3"

	setB := make(map[string]string)
	setB["key2"] = "value4"
	setB["key3"] = "value5"
	setB["key4"] = "value6"

	expected := make(map[string]string)
	expected["key1"] = "value1"
	expected["key2"] = "value4"
	expected["key3"] = "value5"
	expected["key4"] = "value6"

	UpdateFileSecret(setA, setB)

	assert.Equal(t, expected, setB)
}

func TestRunCommand(t *testing.T) {
	setup()
	defer teardown()

	wd, _ := os.Getwd()

	MockCommandsQueue = []CommandSpec{
		{"a b c d 123 '456 ", "ABC\nDEF\n\x33", "GHI\nJKL\n\x33", 100},
		{"e f g h 456 '789 ", "MNO\nPQR\n\x33", "STU\nVWX\n\x33", 101},
		{"i j k l 789 '012 ", "YZA\nBCD\n\x33", "EFG\nHIJ\n\x33", 102},
		{"m n o p 012 '345 ", "KLM\nNOP\n\x33", "QRS\nTUV\n\x33", 0},
		{"q r s t", "", "", 1},
	}

	out, err, code := RunCommandPrintOut(wd, "a", "b", "c", "d", "123", "'456 ")
	assert.Equal(t, "ABC\nDEF\n\x33", out)
	assert.Equal(t, "GHI\nJKL\n\x33", err)
	assert.Equal(t, 123, code)

	out, err, code = RunCommandGetOutput(wd, "e", "f", "g", "h", "456", "'789 ")
	assert.Equal(t, "MNO\nPQR\n\x33", out)
	assert.Equal(t, "STU\nVWX\n\x33", err)
	assert.Equal(t, 456, code)

	code = RunCommandGetStatus(wd, "i", "j", "k", "l", "789", "'012 ")
	assert.Equal(t, 789, code)

	assert.True(t, IsCommandSuccess(wd, "m", "n", "o", "p", "012", "'345 "))

	assert.False(t, IsCommandSuccess(wd, "q", "r", "s", "t"))
}

func TestIsGitRepo(t *testing.T) {
	setup()
	defer teardown()

	wd, _ := os.Getwd()

	MockCommandsQueue = []CommandSpec{
		{"git rev-parse", "", "", 0},
		{"git rev-parse", "", "", 1},
	}

	assert.True(t, IsGitRepo(wd))
	assert.False(t, IsGitRepo(wd))
}

func TestIsProsProject(t *testing.T) {
	setup()
	defer teardown()

	wd, _ := os.Getwd()

	os.WriteFile("project.pros", []byte("any"), 0644)
	assert.True(t, IsProsProject(wd))
	os.Remove("project.pros")
	assert.False(t, IsProsProject(wd))
}

func TestLinkLocalRepoToServerCommand(t *testing.T) {
	setup()
	defer teardown()

	Secret["computer-name"] = "Computer Name"
	Secret["email"] = "BitBucket Email"

	wd, _ := os.Getwd()

	url := GetRepoUrl("some-repo")

	MockCommandsQueue = []CommandSpec{
		// fail to add origin
		{"git rev-parse", "", "", 0},
		{"git remote get-url origin", "", "error: No such remote 'origin'", 1},
		{"git remote add origin " + url, "", "", 1},
		// add origin successfully
		{"git rev-parse", "", "", 0},
		{"git remote get-url origin", "", "error: No such remote 'origin'", 1},
		{"git remote add origin " + url, "", "", 0},
		{"git config user.name Computer Name", "", "", 0},
		{"git config user.email BitBucket Email", "", "", 0},
		{"git config commit.gpgsign false", "", "", 0},
		// set origin successfully
		{"git rev-parse", "", "", 0},
		{"git remote get-url origin", "something else", "", 0},
		{"git remote set-url origin " + url, "", "", 0},
		{"git config user.name Computer Name", "", "", 0},
		{"git config user.email BitBucket Email", "", "", 0},
		{"git config commit.gpgsign false", "", "", 0},
		// set origin successfully
		{"git rev-parse", "", "", 0},
		{"git remote get-url origin", url, "", 0},
		{"git config user.name Computer Name", "", "", 0},
		{"git config user.email BitBucket Email", "", "", 0},
		{"git config commit.gpgsign false", "", "", 0},
	}

	assert.False(t, LinkLocalRepoToServerCommand(wd, "some-repo"))
	assert.True(t, LinkLocalRepoToServerCommand(wd, "some-repo"))
	assert.True(t, LinkLocalRepoToServerCommand(wd, "some-repo"))
	assert.True(t, LinkLocalRepoToServerCommand(wd, "some-repo"))
}

func TestBackupCommand(t *testing.T) {
	setup()
	defer teardown()

	Secret["computer-name"] = "Computer Name"
	Secret["email"] = "BitBucket Email"

	wd, _ := os.Getwd()

	MockCommandsQueue = []CommandSpec{
		// error 105
		{"git rev-parse", "", "", 0},
		{"git add -A", "", "", 0},
		{"git commit -m Backup", "", "", 1},
		// error 106
		{"git rev-parse", "", "", 0},
		{"git add -A", "", "", 0},
		{"git commit -m Backup", "", "", 0},
		{"git push -u origin master", "", "", 1},
		// success
		{"git rev-parse", "", "", 0},
		{"git add -A", "", "", 0},
		{"git commit -m Backup", "", "", 0},
		{"git push -u origin master", "", "", 0},
	}

	assert.False(t, BackupCommand(wd))
	assert.False(t, BackupCommand(wd))
	assert.True(t, BackupCommand(wd))
}

func TestBeep(t *testing.T) {
	BeepSuccess()
	BeepFail()
}
