package config

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type netConf struct {
	Ip    	[4]byte
	MAC	   	[6]byte
	Judge   bool
	MTU     int
	Time 	time.Duration
}

type Configs struct {
	Cfgs netConf
}

//根据路径初始化配置文件
func configRead(path string) map[string]string{
	confMap := make(map[string]string)

	confFile, err := os.Open(path)
	defer confFile.Close()
	if err != nil{
		log.Fatal("Error : conf file open failed :",err)
	}
	tmpBuffer := bufio.NewReader(confFile)
	for{
		lineRaw, _, err := tmpBuffer.ReadLine()
		if err != nil{
			//文件读取结束，退出循环
			if err == io.EOF{
				break
			}
			log.Fatal("Error : conf file read lines :", err)
		}
		//去除空格并转换为字符串
		lineStr := strings.TrimSpace(string(lineRaw))
		//过滤以 #   ; 开头的注释
		if len(lineStr) == 0 || lineStr[0] == '#' || lineStr[0] == ';' {
			continue
		}
		//log.Print(lineStr)
		//找到=在字符串中的位置
		equalIndex := strings.Index(lineStr, "=")
		if equalIndex < 0{
			log.Fatal("Error : conf file the index of '=' < 0")
			break
		}
		//获取key
		key := strings.TrimSpace(lineStr[:equalIndex])
		if (len(key) <= 0){
			log.Fatal("Error : conf file len(key) <= 0")
			break
		}
		//获取value
		value := strings.TrimSpace(lineStr[equalIndex + 1:])
		if (len(value) <= 0){
			log.Fatal("Error : conf file len(value) <= 0")
			break
		}
		confMap[key] = value
	}
	return confMap
}
// init config
func InitConfigs() Configs{

	netCfg := configRead("./config/cfg.conf")
	if len(netCfg) == 0{
		log.Fatal("Error : conf file read failed")
	}
	return typeTrans(netCfg)
}

func typeTrans(cfg map[string]string) Configs{
	var configs Configs
	configs.Cfgs.Ip = strtoIp(cfg["Ip"])
	configs.Cfgs.MTU = strtoInt(cfg["MTU"])
	configs.Cfgs.Judge = strtoBool(cfg["Judge"])
	configs.Cfgs.MAC = strtoMac(cfg["MAC"])
	configs.Cfgs.Time = strtoTime(cfg["Time"])
	return configs
}

func strtoInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil{
		log.Fatal("[error] : cfg str to Int ",err)
	}
	return value
}

func strtoIp(str string) [4]byte{
	var ip [4]byte
	ipStr := strings.Split(str, ".")
	if len(ipStr) != 4{
		log.Fatal("ERROR: please check the conf ip address.")
	}
	ip[0] = byte(strtoInt(ipStr[0]))
	ip[1] = byte(strtoInt(ipStr[1]))
	ip[2] = byte(strtoInt(ipStr[2]))
	ip[3] = byte(strtoInt(ipStr[3]))
	return ip
}

func strtoMac(str string) [6]byte{
	var mac [6]byte
	macStr := strings.Split(str, ":")
	if len(macStr) != 6{
		log.Fatal("ERROR: please check the conf ip address.")
	}
	for i:=0; i < 6; i++{
		tmp, _ := strconv.ParseUint("0x"+macStr[i], 0, 0)
		mac[i] = byte(tmp)
	}
	return mac
}

func strtoTime(str string) time.Duration{
	var t time.Duration
	strLen := len(str)
	timeUnit := str[strLen-1:]
	timeNum := strtoInt(str[:strLen - 1])
	if timeUnit == "Y"{
		t =  12 * 30 * 24 * time.Hour
	}else if timeUnit == "M"{
		t =  30 * 24 * time.Hour
	}else if timeUnit == "D"{
		t =  24 * time.Hour
	}
	for timeNum > 1{
		t += t
		timeNum--
	}
	return t
}

func strtoBool(str string) bool {
	if str == "false"{
		return false
	}else{
		return true
	}
}

/* Byte转Ip */
func ByteToIp(num [4]byte) string {
	return  strconv.Itoa(int(num[0]))+"."+ strconv.Itoa(int(num[1]))+"."+ strconv.Itoa(int(num[2]))+"."+ strconv.Itoa(int(num[3]))
}
