package main


import (
    "fmt"
    "time"
    "crypto/sha512"
)


type Block interface{
    bHash() []byte
}


type RawBlock struct{
    index int
    timestamp int64
    data string
    previousHash []byte
}

func printRawBlock(b RawBlock) {
    fmt.Printf("index: %d, timestamp: %d, data: %s, previousHash: %x\n", b.index, b.timestamp, b.data, b.previousHash)
}


func (b RawBlock) bHash() []byte{
    hashBlock := sha512.New()
    hashBlock.Write([]byte(fmt.Sprintf("%v", b)))
    return hashBlock.Sum(nil)
}


func getBlockHash(b Block) []byte{
    return b.bHash()
}


func main() {
    genesis := RawBlock{0, time.Now().Unix(), "This is the genesis block", []byte{0}}

    currentBlock := genesis

    for i := 0; i < 20; i++{

        printRawBlock(currentBlock)
        fmt.Printf("sha512:\t%x\n", getBlockHash(currentBlock))

        time.Sleep(time.Millisecond * 1000)

        currentBlock = RawBlock{i, time.Now().Unix(), fmt.Sprintf("This is the %d block", i), getBlockHash(currentBlock)}
    }

    printRawBlock(currentBlock)
    fmt.Printf("sha512:\t%x\n", getBlockHash(currentBlock))
}
