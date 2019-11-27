/**
 * @Author: DollarKillerX
 * @Description: 本脚本dns实践
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午8:26 2019/11/27
 */
package test

import (
	"fmt"
	dns2 "github.com/Q2h1Cg/dnsbrute/dns"
	"github.com/dollarkillerx/easyutils/clog"
	"github.com/miekg/dns"
	"log"
	"testing"
	"time"
)

func TestDnsM(t *testing.T) {
	domains := []string{
		"www.dollarkiller.com",
		"dollarkiller.com",
		"ps.cs",
		"xxxp.baidu.com",
	}

	for _, domain := range domains {
		testDomain(domain)
	}
}

func testDomain(domain string) {
	defer func() {
		log.Println("===============")
		fmt.Println()
		fmt.Println()
	}()
	conn, e := newDns("8.8.8.8:53")
	if e != nil {
		panic(e)
	}

	// 进行dns查询
	msg := &dns.Msg{}
	// 进行A记录的查询
	msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	if err := conn.SetWriteDeadline(time.Now().Add(time.Second)); err != nil {
		panic(err)
	}

	if err := conn.WriteMsg(msg); err != nil {
		panic(err)
	}
	log.Println(domain)
	var err error
	if err := conn.SetReadDeadline(time.Now().Add(time.Second)); err != nil {
		panic(err)
	}
	if msg, err = conn.ReadMsg(); err != nil || len(msg.Question) == 0 {
		panic(err)
	}

	record := dns2.NewRecord(domain, msg.Answer)
	if record == nil {
		clog.PrintEr("eee")
		return
	}
	clog.Println(record)
}

func newDns(host string) (*dns.Conn, error) {
	return dns.DialTimeout("udp", host, time.Second)
}
