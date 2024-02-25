package microBatchService

import (
	"net/http"

	"github.com/Dhruv1993/app"
	"github.com/go-chi/render"
)

// @This api handle money transfers between one account to another
// @Summary Transfer money between accounts
// @Description Process a money transfer between accounts
// @ID transfer
// @Accept json
// @Produce json
// @Param 200 body TransferRequest true "Transfer request"
// @Success 200 {object} TransferResponse "Transfer processed successfully"
// @Router /api/v1/transfer [post]
func (s Service) TransferHandler() app.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var job TransferRequest
		
		err := app.DecodeJSON(r.Body, &job)
		if err != nil {
			return err
		}

		// Check if any of the required fields in TransferRequest is empty
		if job.FromAccount == "" || job.ToAccount == "" || job.Amount <= 0 {
			return app.BadRequest(err)
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

func HeartbeatHandler(w http.ResponseWriter, r *http.Request) app.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("."))

			return nil
	}
	
}