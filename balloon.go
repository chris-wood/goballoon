package main

import "os"
import "fmt"
import "time"
import "strconv"
import "bytes"
import "runtime"
import "encoding/binary"
import "github.com/codahale/blake2"

func hash(count int, values ...[]byte) []byte {
    h := blake2.NewBlake2B()

    bs := make([]byte, 4)
    binary.LittleEndian.PutUint32(bs, 31415926)
    h.Write(bs)

    for _, block := range values {
        h.Write(block)
    }

    return h.Sum(nil)
}

func balloon(password, salt string, s_cost, t_cost int) []byte {
    delta := 3
    count := 0

    blocks := [][]byte{}

    // Expansion loop
    block := hash(count, []byte(password), []byte(salt))
    count += 1
    blocks = append(blocks, block)
    for i := 1; i < s_cost; i++ {
        block = hash(count, blocks[i - 1])
        count += 1
        blocks = append(blocks, block)
    }

    // Mixing loop
    for t := 0; t < t_cost; t++ {
        for m := 0; m < s_cost; m++ {
            index := ((m + s_cost) - 1) % s_cost
            prev := blocks[index]
            blocks[m] = hash(count, prev, blocks[m])
            count += 1

            // Delta pseudorandom mixing
            for i := 0; i < delta; i++ {
                block = blocks[(t * m * i) % s_cost] // poor man's int_to_block -- not pseudorandom
                hval := hash(count, []byte(salt), block)
                count += 1

                // Convert the hash to an integer and use it as the index of the next mixin
                buf := bytes.NewBuffer(hval)
                other, _ := binary.ReadVarint(buf)
                index := (int(other % int64(s_cost)) + s_cost) % s_cost

                blocks[m] = hash(count, blocks[m], blocks[index])
                count += 1
            }
        }
    }

    // Extraction
    return blocks[s_cost - 1]
}

func main() {
    if len(os.Args) != 5 {
        fmt.Println("usage: balloon <password> <salt> <n> <r>")
        os.Exit(-1)
    }

    p := os.Args[1]
    s := os.Args[2]
    n, _ := strconv.Atoi(os.Args[3])
    r, _ := strconv.Atoi(os.Args[4])

    // Compute the hash, but swallow the output
    start := time.Now()
    balloon(p, s, n, r) 
    delta := time.Since(start)

    // Output format: n, r, runtime, allocation, total_allocation
    var mem runtime.MemStats
    runtime.ReadMemStats(&mem)
    fmt.Printf("%d,%d,%s,%d,%d\n", n, r, delta, mem.Alloc, mem.TotalAlloc)
}
