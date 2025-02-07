package slugs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/opensea"
	"github.com/benleb/gloomberg/internal/rueidica"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/ethereum/go-ethereum/common"
)

// ErrFetchingBlurSlug is returned when the blur slug could not be fetched.
var ErrFetchingBlurSlug = fmt.Errorf("error fetching blur slug")

type BlurSlugResponse struct {
	Success     bool `json:"success"`
	Collections []struct {
		ContractAddress string `json:"contractAddress"`
		CollectionSlug  string `json:"collectionSlug"`
		Name            string `json:"name"`
		ImageURL        string `json:"imageUrl"`
		TotalSupply     int    `json:"totalSupply"`
	} `json:"collections"`
}

func SlugWorker(slugTicker *time.Ticker, slugQueue *chan common.Address, rueidica *rueidica.Rueidica) {
	for address := range *slugQueue {
		gbl.Log.Infof("fetching opensea slug for: %s", address.Hex())

		if osslug := opensea.GetCollectionSlug(address); osslug != "" {
			err := rueidica.StoreOSSlug(context.Background(), address, osslug)
			if err != nil {
				gbl.Log.Errorf("error storing opensea slug: %s | %s", style.ShortenAddress(&address), err)
			}

			gbl.Log.Infof("caching opensea slug for: %s\n", address.Hex())
		} else {
			gbl.Log.Warnf("❌ fetching opensea slug for collection %s failed: no slug found", address.Hex())

			return
		}

		gbl.Log.Infof("fetching blur slug for: %s", address.Hex())

		if blurslug, err := GetBlurSlugByName(address); err == nil && blurslug != "" {
			err := rueidica.StoreBlurSlug(context.Background(), address, blurslug)
			if err != nil {
				gbl.Log.Errorf("error storing blur slug: %s | %s", style.ShortenAddress(&address), err)
			}

			gbl.Log.Infof("caching blur slug for: %s\n", address.Hex())
		} else {
			gbl.Log.Warnf("❌ fetching blur slug for collection %s failed: no slug found", address.Hex())

			return
		}

		<-slugTicker.C
	}
}

func GetBlurSlugByName(collectionAddress common.Address) (string, error) {
	// build url
	url := "https://core-api.prod.blur.io/v1/search?query=" + collectionAddress.String()

	response, err := utils.HTTP.Get(context.Background(), url)
	if err != nil {
		if os.IsTimeout(err) {
			gbl.Log.Warnf("⌛️ timeout while fetching blur slug: %+v", err.Error())
		} else {
			gbl.Log.Errorf("❌ blur slug error: %+v", err.Error())
		}

		return "", err
	}

	gbl.Log.Infof("blur slug response status: %s", response.Status)

	defer response.Body.Close()

	if collectionList, err := parseBlurSlug(response); err == nil && collectionList != nil && len(collectionList.Collections) > 0 {
		gbl.Log.Infof("blur slug: %s", collectionList.Collections[0].CollectionSlug)

		return collectionList.Collections[0].CollectionSlug, nil
	}

	return "", err
}

func parseBlurSlug(response *http.Response) (*BlurSlugResponse, error) {
	bodyBytes, err := io.ReadAll(response.Body)

	defer func() { _ = response.Body.Close() }()

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, ErrFetchingBlurSlug
	}

	var blurSlugResponse BlurSlugResponse

	err = json.Unmarshal(bodyBytes, &blurSlugResponse)
	if err != nil {
		return nil, err
	}

	return &blurSlugResponse, nil
}
