package microBatchService

import (
	"net/http"

	"github.com/Dhruv1993/app"
	"github.com/go-chi/render"
)

func (s Service) TaskHandler() app.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var job TransferRequest
		err := app.DecodeJSON(r.Body, &job)
		if err != nil {
			return err
		}

		tj := Transfer{
			FromAccount: job.FromAccount,
			ToAccount:   job.ToAccount,
			Amount:      job.Amount,
		}
		d, err := s.API.processTransfers(tj)
		if err != nil {
			return err
		}

		//Respond with a success message and the list of tasks
		render.JSON(w, r, TransferResponse{
			Message:  "Transfer processed successfully",
			Transfer: d,
			status:   "OK",
		})
		return nil
	}
}
