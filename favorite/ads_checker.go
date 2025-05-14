package favorite

type AdsChecker interface {
	IsAdsExists(title string) bool
}
