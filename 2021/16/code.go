package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
)

type Packet interface {
	packetVersion() int64
	packetType() int64
}

type PacketData struct {
	Value *big.Int
	Size  uint
}

type decodedPacketData struct {
	pVersion int64
	pType    int64
	pValue   *big.Int
}

type operatorPacket struct {
	pVersion int64
	pType    int64
	Packets  []Packet
}

// for the Packet interface
func (packet decodedPacketData) packetVersion() int64 {
	return packet.pVersion
}

func (packet decodedPacketData) packetType() int64 {
	return packet.pType
}

func (operator operatorPacket) packetVersion() int64 {
	return operator.pVersion
}

func (operator operatorPacket) packetType() int64 {
	return operator.pType
}

func (packet PacketData) extractBits(startBit uint, length uint) *big.Int {
	tmp := new(big.Int)
	mask := new(big.Int)
	one := big.NewInt(1)

	mask.Sub(mask.Lsh(one, length), one)

	tmp.And(tmp.Rsh(packet.Value, packet.Size-startBit-length), mask)

	return tmp
}

func parseArgs() (inputFile string) {
	if len(os.Args) == 2 {
		inputFile = os.Args[1]
	} else {
		log.Fatalf("Usage: %v INPUTFILE", os.Args[0])
	}

	return
}

func getData(inputFile string) (inputData []string) {
	fh, err := os.Open(inputFile)

	if err != nil {
		log.Fatalln(err)
	}

	fileScanner := bufio.NewScanner(fh)

	for fileScanner.Scan() {
		inputData = append(inputData, fileScanner.Text())
	}

	return
}

func parsePacket(packet PacketData, startBit uint) (Packet, uint) {
	pVersion := packet.extractBits(startBit, 3).Int64()
	pType := packet.extractBits(startBit+3, 3).Int64()
	startBit += 6

	if pType == 4 {
		// literal value
		// 110 100 10111 11110 00101 000
		// VVV TTT AAAAA BBBBB CCCCC ---
		pVal := big.NewInt(0)

		for {
			// loop over them bits
			tmp := packet.extractBits(startBit, 5).Int64()
			startBit += 5

			if tmp >= 16 {
				pVal.Add(pVal, big.NewInt(tmp&15))
				pVal.Lsh(pVal, 4)
			} else {
				pVal.Add(pVal, big.NewInt(tmp))
				break
			}

		}
		return decodedPacketData{pVersion: pVersion, pType: pType, pValue: pVal}, startBit

	} else {
		// operator packet
		// 001 110 0 000000000011011 11010001010 0101001000100100 0000000
		// VVV TTT I LLLLLLLLLLLLLLL AAAAAAAAAAA BBBBBBBBBBBBBBBB -------

		embeddedPackets := make([]Packet, 0)
		operatorPackets := operatorPacket{pVersion: pVersion, pType: pType, Packets: embeddedPackets}

		lengthTypeId := packet.extractBits(startBit, 1).Uint64()
		startBit++

		if lengthTypeId == 0 {
			// 15 bit number representing nr of bits
			tmp := packet.extractBits(startBit, 15).Int64()
			startBit += 15
			endBit := startBit + uint(tmp)

			for startBit < endBit {
				var tmp2 Packet
				tmp2, startBit = parsePacket(packet, startBit)
				operatorPackets.Packets = append(operatorPackets.Packets, tmp2)
			}
		} else if lengthTypeId == 1 {
			// 11 bit number representing nr of subpackets
			packetsToExtract := packet.extractBits(startBit, 11).Int64()
			startBit += 11

			for i := int64(0); i < packetsToExtract; i++ {
				var tmp2 Packet
				tmp2, startBit = parsePacket(packet, startBit)
				operatorPackets.Packets = append(operatorPackets.Packets, tmp2)
			}
		}

		return operatorPackets, startBit
	}
}

func convertHex(inputData string) PacketData {
	tmp := new(big.Int)
	tmp.SetString(inputData, 16)

	packet := PacketData{Value: tmp, Size: uint(len(inputData) * 4)}

	return packet
}

func getVersionSum(packet Packet) (result int64) {
	result += packet.packetVersion()

	if operator, ok := packet.(operatorPacket); ok {
		// whee it's an operator
		for _, embeddedPacket := range operator.Packets {
			result += getVersionSum(embeddedPacket)
		}
	}

	return
}

func calculateBITSValue(packet Packet) (result *big.Int) {
	result = big.NewInt(0)

	switch aPacket := packet.(type) {
	case decodedPacketData:
		return aPacket.pValue

	case operatorPacket:
		switch aPacket.packetType() {
		case 0:
			// sum
			for _, subPacket := range aPacket.Packets {
				result.Add(result, calculateBITSValue(subPacket))
			}

		case 1:
			// product
			// start with 1, because otherwise...
			result = big.NewInt(1)
			for _, subPacket := range aPacket.Packets {
				result.Mul(result, calculateBITSValue(subPacket))
			}

		case 2:
			// min
			result = big.NewInt(math.MaxInt)
			for _, subPacket := range aPacket.Packets {
				tmp2 := calculateBITSValue(subPacket)
				if result.Cmp(tmp2) > 0 {
					result = tmp2
				}
			}

		case 3:
			// max
			result = big.NewInt(math.MinInt)
			for _, subPacket := range aPacket.Packets {
				tmp2 := calculateBITSValue(subPacket)
				if result.Cmp(tmp2) < 0 {
					result = tmp2
				}
			}

		case 5:
			// greater than, 2 subpackets
			if calculateBITSValue(aPacket.Packets[0]).Cmp(calculateBITSValue(aPacket.Packets[1])) > 0 {
				result = big.NewInt(1)
			}

		case 6:
			// less then, 2 subpackets
			if calculateBITSValue(aPacket.Packets[0]).Cmp(calculateBITSValue(aPacket.Packets[1])) < 0 {
				result = big.NewInt(1)
			}
		case 7:
			// equals, 2 subpackets
			if calculateBITSValue(aPacket.Packets[0]).Cmp(calculateBITSValue(aPacket.Packets[1])) == 0 {
				result = big.NewInt(1)
			}
		}
	}

	return
}

func main() {
	inputFile := parseArgs()
	inputData := getData(inputFile)

	packet := convertHex(inputData[0])
	parsedPackets, _ := parsePacket(packet, 0)

	result1 := getVersionSum(parsedPackets)
	fmt.Printf("Part 1: total versions: %d\n", result1)

	result2 := calculateBITSValue(parsedPackets)
	fmt.Printf("Part 2: expresssion value: %d\n", result2)

}
