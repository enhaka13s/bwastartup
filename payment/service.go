package payment

import (
	midtrans "github.com/veritrans/go-midtrans"
	"bwastartup/user"
	"strconv"
)

type service struct {
}

type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error) {
	midclient := midtrans.NewClient()
    midclient.ServerKey = "SB-Mid-server-JuyWzhYNRQF91mZYOaXUytE0"
    midclient.ClientKey = "SB-Mid-client-nuRtH-K-RFoBd7Lf"
    midclient.APIEnvType = midtrans.Sandbox

    snapGateway := midtrans.SnapGateway{
        Client: midclient,     
    }

	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}

