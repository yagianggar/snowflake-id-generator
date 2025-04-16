package id

import (
	"errors"
	"fmt"
	"hash/fnv"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/sony/sonyflake"
)

type (
	sonyflakeGenerator struct {
		*sonyflake.Sonyflake
	}
)

var (
	sonyFlakeGenerator *sonyflakeGenerator = nil
	DefaultStartTime                       = time.Date(2023, 1, 1, 0, 0, 0, 0, GetLocalTime())
)

// SonyFlakeGenerator creates a new id generator
// the function panics if the generator cannot be created
func SonyFlakeGenerator() *sonyflakeGenerator {
	if sonyFlakeGenerator == nil {
		sfg := &sonyflakeGenerator{
			Sonyflake: sonyflake.NewSonyflake(sonyflake.Settings{
				MachineID: hostname,
				StartTime: DefaultStartTime,
			}),
		}

		sonyFlakeGenerator = sfg
	}

	return sonyFlakeGenerator
}

func (s *sonyflakeGenerator) NextID() (uint64, error) {
	return s.Sonyflake.NextID()
}

func (s *sonyflakeGenerator) GenerateID() (string, error) {
	id, err := s.NextID()
	if err != nil {
		return "", err
	}

	return strconv.FormatUint(id, 10), nil
}

func hostname() (uint16, error) {
	host, err := os.Hostname()
	if err != nil {
		return 0, err
	}

	h := fnv.New32()
	_, hashErr := h.Write([]byte(host))
	if hashErr != nil {
		return 0, hashErr
	}

	return uint16(h.Sum32()), nil
}

func GetLocalTime() *time.Location {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return loc
}

const ELECTRUM_PREFIX = "EL"
const ORDER_CATEGORY = "MANUAL"

func GetOrderNumber() (string, error) {
	idStr, err := SonyFlakeGenerator().GenerateID()
	if err != nil {
		return "", fmt.Errorf("error bro")
	}

	id := fmt.Sprintf("%s-%s-%s", ELECTRUM_PREFIX, ORDER_CATEGORY, idStr)

	return id, nil
}

func ExtractIdInfo(orderNumber string) (uint64, time.Time, error) {
	idStr, err := ValidateAndExtractSnowflakeID(orderNumber)
	if err != nil {
		return 0, time.Time{}, err
	}

	return ExtractID(idStr)
}

func ValidateAndExtractSnowflakeID(orderNumber string) (string, error) {
	r := regexp.MustCompile(`^EL-([A-Z0-9\-]+)-(\d+)$`)
	matches := r.FindStringSubmatch(orderNumber)

	if len(matches) != 3 {
		return "", errors.New("invalid order number format")
	}

	return matches[2], nil
}
