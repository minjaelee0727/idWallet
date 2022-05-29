package blockchain

import (
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/minjaelee0727/idWallet/backend/constant"
	"github.com/minjaelee0727/idWallet/backend/db"
	"github.com/minjaelee0727/idWallet/backend/utils"
	"github.com/minjaelee0727/idWallet/backend/wallet"
)

const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5
	blockInterval      int = 2
	allowedRange       int = 2
)

type blockchain struct {
	NewestHash        string `json:"newestHash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:"currentDifficulty"`
}

type Block struct {
	Hash       string            `json:"hash"`
	PrevHash   string            `json:"prevHash,omitempty"`
	Height     int               `json:"height"`
	Difficulty int               `json:"difficulty"`
	Nonce      int               `json:"nonce"`
	Timestamp  int               `json:"timestamp"`
	Credential CredentialOnBlock `json:"credential"`
}

type CredentialOnBlock struct {
	Id        string `json:"CdId"`        // hash value of this credential
	Timestamp int    `json:"CdTimestamp"` // just for hash
	Active    bool   `json:"CdActive"`    // check whether it is active
	Signature string `json:"CdSignature"` // Block will save signature only to confirm that idWallet has not been changed
	// Only user have personal info in their idWallet
	// Signature is created by Credential + private key
}

var b *blockchain

func Blocks(b *blockchain) []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, errors.New("block not found")
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func recalculateDifficulty(b *blockchain) int {
	allBlocks := Blocks(b)
	newestBlock := allBlocks[0]
	lastRecalculatedBlock := allBlocks[difficultyInterval-1]
	actualTime := (newestBlock.Timestamp / 60) - (lastRecalculatedBlock.Timestamp / 60)
	expectedTime := difficultyInterval * blockInterval
	if actualTime <= (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime >= (expectedTime + allowedRange) {
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}

func getDifficulty(b *blockchain) int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return recalculateDifficulty(b)
	} else {
		return b.CurrentDifficulty
	}
}

func (b *blockchain) AddBlock(cr constant.Credential, w *wallet.IdWallet) {
	block := createBlock(b.NewestHash, b.Height+1, getDifficulty(b))
	cd := &CredentialOnBlock{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		Active:    true,
		Signature: "",
	}
	cd.Id = utils.Hash(cd)
	hashedCredential := utils.Hash(cr)
	cd.Signature = wallet.MakeSignature(w, hashedCredential)
	block.Credential = *cd
	persistBlock(block)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	persistBlockhain(b)
}

func (b *blockchain) AddGenesis() {
	block := createBlock(b.NewestHash, b.Height+1, getDifficulty(b))
	persistBlock(block)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	persistBlockhain(b)
}

func persistBlockhain(b *blockchain) {
	db.SaveCheckpoint(utils.ToBytes(b))
}

var once sync.Once

func Blockchain() *blockchain {
	once.Do(func() {
		b = &blockchain{
			Height: 0,
		}
		checkpoint := db.Restart()
		if checkpoint == nil {
			b.AddGenesis()
		} else {
			b.restore(checkpoint)
		}
	})
	return b
}

// 1 block per 1 credential for now
func createBlock(prevHash string, height, diff int) *Block {
	block := &Block{
		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: diff,
		Nonce:      0,
	}
	block.mine()
	return block
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func persistBlock(b *Block) {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}
