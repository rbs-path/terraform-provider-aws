package finder

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/servicecatalog"
)

func PortfolioShare(conn *servicecatalog.ServiceCatalog, portfolioID, shareType, accountID, orgNodeValue string) (*servicecatalog.PortfolioShareDetail, error) {
	input := &servicecatalog.DescribePortfolioSharesInput{
		PortfolioId: aws.String(portfolioID),
		Type:        aws.String(shareType),
	}
	var result *servicecatalog.PortfolioShareDetail

	err := conn.DescribePortfolioSharesPages(input, func(page *servicecatalog.DescribePortfolioSharesOutput, lastPage bool) bool {
		if page == nil {
			return !lastPage
		}

		for _, deet := range page.PortfolioShareDetails {
			if deet == nil {
				continue
			}

			if aws.StringValue(deet.PrincipalId) == accountID && accountID != "" {
				result = deet
				return false
			}

			if aws.StringValue(deet.PrincipalId) == orgNodeValue && orgNodeValue != "" {
				result = deet
				return false
			}
		}

		return !lastPage
	})

	return result, err
}
