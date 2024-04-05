package main

import (
	"encoding/json"
	"fmt"
	"hackattic/utils"
	"os"
	"os/exec"
	"sync"

	"golang.org/x/crypto/ssh"

	l "github.com/charmbracelet/log"
)

const chalName = "a_global_presence"

var log *l.Logger

type chalBody struct {
	Ptoken string `json:"presence_token"`
}

type solBody struct {
}

func init() {
	log = utils.GetLogger(chalName)
}

type connHandler struct {
	host     string
	pToken   string
	userName string
	location string
	session  *ssh.Session
	reqUrl   string
	wg       *sync.WaitGroup
	cmd      string
}

func (c *connHandler) createSession() {
	pkeyPath := "/home/aryan/.ssh/id_ed25519"
	var authMethod []ssh.AuthMethod
	pkey, err := os.ReadFile(pkeyPath)
	if err != nil {
		log.Fatal(err)
	}
	signer, err := ssh.ParsePrivateKey(pkey)
	authMethod = []ssh.AuthMethod{ssh.PublicKeys(signer)}

	config := &ssh.ClientConfig{
		User:            c.userName,
		Auth:            authMethod,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", c.host), config)
	if err != nil {
		log.Fatalf("Failed to create client for %s : %s", c.location, err)
	}
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session for %s : %s", c.location, err)
	}
	c.reqUrl = fmt.Sprintf("https://hackattic.com/_/presence/$%s", c.pToken)

	c.cmd = fmt.Sprintf("curl -s https://hackattic.com/_/presence/$%s", c.pToken)
    // establish presence
	log.Debugf("Command for %s : %s", c.location, c.cmd)
	output, err := session.CombinedOutput(c.cmd)
	log.Infof("%s : %s", c.location, string(output))
	// err := c.session.Run(cmd)
	if err != nil {
		log.Fatalf("Err curling at %s: %s", c.location, err)
	}
	c.wg.Done()
}

func (c *connHandler) establishPresence(s *ssh.Session) {
	// cmd := fmt.Sprintf("curl %s", c.reqUrl)
	log.Debugf("Command for %s : %s", c.location, c.cmd)
	output, err := s.CombinedOutput(c.cmd)
	log.Infof("%s : %s", c.location, string(output))
	// err := c.session.Run(cmd)
	if err != nil {
		log.Fatalf("Err curling at %s: %s", c.location, err)
	}
	c.wg.Done()
}

func hiFromIndia(wg *sync.WaitGroup, pToken string) {
	curlUrl := fmt.Sprintf("https://hackattic.com/_/presence/$%s", pToken)
	log.Debug(curlUrl)
	cmd := exec.Command("curl", "-s", curlUrl)
	out, err := cmd.CombinedOutput()
	log.Infof("India : %s", string(out))
	if err != nil {
		log.Fatalf("Err curling at %s: %s", "india", err)

	}
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(7)
	b := chalBody{}
	s := solBody{}
	resp := utils.GetChal(chalName)
	log.Debug(string(resp))
	json.Unmarshal(resp, &b)
	log.Debug(b)

	var droplets []connHandler
	locations := []string{"banglore", "frankfurt", "amsterdam", "sydney", "singapore", "toronto", "newYork"}
	ipAddrs := []string{"142.93.210.167", "139.59.156.58", "165.232.92.187", "170.64.196.171", "209.97.166.127", "142.93.149.190", "167.172.225.168"} // add the list of ip addresses here

	for idx, loc := range locations {
		vps := connHandler{}
		vps.location = loc
		vps.host = ipAddrs[idx]
		vps.userName = "root"
		vps.pToken = b.Ptoken
		vps.wg = &wg
		droplets = append(droplets, vps)
	}

	// go hiFromIndia(&wg, b.Ptoken)
	for _, vps := range droplets {
		go vps.createSession()
	}

	wg.Wait()

	jsonData, err := json.Marshal(s)
	if err != nil {
		log.Errorf("Err marshalling %v", err)
	}
	status := utils.SendSol(chalName, jsonData)
	log.Info("Status: %v\n", status)
}
