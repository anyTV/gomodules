package idgen

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"

	logger "github.com/anyTV/gomodules/v2/logging"
	"github.com/sixafter/nanoid"
)

func GenerateId(prefix string) (string, error) {
	rand, err := IDGenerator.New()
	if err != nil {
		return "", errors.Join(err, errors.New("error generating id"))
	}

	return prefix + rand.String(), nil
}

func Generate() string {
	id, err := GenerateId("")
	if err != nil {
		logger.Warnf("id.Generate() failed to generate, err: %v", err)
	}

	return id
}

const HeartbeatPrefix = "he"

// GenerateCustomId
//
// Creates a heartbeat id with the format: heXXXXXXX
func GenerateCustomId() (string, error) {
	return GenerateId(HeartbeatPrefix)
}

func StripNonAlpha(s string) string {
	var b strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) {
			b.WriteRune(r)
		}
	}

	return b.String()
}

func CreateRandomHandle() string {
	// TODO: implement handle generation
	return ""
}

var HandleLength = 10

// PadWithSuffix
//
// pad the end of the username with numbers
func PadSuffix(s string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for len(s) < HandleLength {
		s += strconv.Itoa(r.Intn(10))
	}

	return s
}

// CreateSuffixNumber
func CreateRandomNumber(size int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var s = ""

	for len(s) < size {
		s += strconv.Itoa(r.Intn(10))
	}

	return s
}

func GenerateHandle(givenName, familyName string) string {
	var initialHandle = StripNonAlpha(givenName) + StripNonAlpha(familyName)

	if len(initialHandle) < 5 {
		initialHandle = CreateRandomHandle()
	}

	initialHandle += CreateRandomNumber(5)

	return PadSuffix(initialHandle)
}

var IDCharSet = "abcdefghijklmnopqrstuvwxyz0123456789"
var IDLength = 7
var IDGenerator nanoid.Interface

func Initialize() error {
	var err error

	IDGenerator, err = nanoid.NewGenerator(
		nanoid.WithAlphabet(IDCharSet),
		nanoid.WithLengthHint(uint16(IDLength)),
	)

	return err
}
