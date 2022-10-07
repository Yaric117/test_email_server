package v1

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"testproject/helpers"
	"testproject/repository"
	"time"
)

const loggerApiTitle = "[api][v1]"

// HandleTest
// @Summary     test
// @Description test
// @Tags        test
// @Accept   json
// @Produce  json
// @Success     200 {object} dto.Response
// @Failure  422          {object} dto.Response
// @Failure     404 {object} dto.Response
// @Router      /test [get]
// @Security    ApiKeyAuth
func HandleTest(w http.ResponseWriter, r *http.Request) {
	//logSectionTitle := loggerApiTitle + "[HandleGetChannelsByUser]"

	var payload interface{}

	RespondSuccess(w, "", payload)
	return
}

// HandleCreateData
// @Summary  Create objects and windows
// @Tags     objects
// @Accept   json
// @Produce  json
// @Param    countObjects body     string true "Count objects"
// @Param    countWindows body     string true "Count object`s windows"
// @Success  200          {object} dto.Response{payload=ResponseCounts}
// @Failure     422 {object} dto.Response
// @Failure  400          {object} dto.Response
// @Router   /create [post]
// @Security ApiKeyAuth
func HandleCreateData(w http.ResponseWriter, r *http.Request) {
	var payload interface{}
	logSectionTitle := loggerApiTitle + "[HandleCreateData]"
	requestData := RequestCounts{}
	ctx := context.Background()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)

	if err != nil {
		RespondError(w, "Data read error!", payload, http.StatusBadRequest)
		log.Printf("%s request body decode failed: %s\n", logSectionTitle, err)
		return
	}

	if requestData.CountObjects == "" {
		RespondError(w, "Object`s count is required!", payload, http.StatusUnprocessableEntity)
		return
	}

	if requestData.CountWindows == "" {
		RespondError(w, "Window`s count is required!", payload, http.StatusUnprocessableEntity)
		return
	}

	countObject, err := strconv.Atoi(requestData.CountObjects)

	if err != nil {
		RespondError(w, "Object`s count is not integer!", payload, http.StatusUnprocessableEntity)
		return
	}

	countWindows, err := strconv.Atoi(requestData.CountWindows)

	if err != nil {
		RespondError(w, "Window`s count is not integer!", payload, http.StatusUnprocessableEntity)
		return
	}

	responseData := ResponseCounts{}

	if countObject != 0 {
		err = repository.DeleteAllObjects(ctx)

		if err != nil {
			RespondError(w, "Delete objects error!", payload, http.StatusBadRequest)
			return
		}

		for i := 1; i <= countObject; i++ {
			object := repository.Object{
				Name:     "Object" + strconv.Itoa(i),
				Timezone: i,
			}
			id, err := object.Insert(ctx)

			if err != nil {
				log.Printf("%s %s", logSectionTitle, err)
				responseData.ErrorObjects += 1
				responseData.ErrorWindows += countWindows
				continue
			}

			responseData.CountObjects += 1

			if countWindows != 0 {
				object.Id = id
				intervalWindow := helpers.RandInterval(2022)

				for j := 1; j <= countWindows; j++ {
					window := repository.ObjectWindow{
						ObjectId:  object.Id,
						DateStart: intervalWindow.Start,
						DateEnd:   intervalWindow.End,
					}
					_, err := window.Insert(ctx)

					if err != nil {
						responseData.ErrorWindows += 1
						continue
					}
					responseData.CountWindows += 1

					intervalWindow.Start = intervalWindow.End
					delta := rand.Intn(86400) + 1
					intervalWindow.End = intervalWindow.Start.Add(time.Second * time.Duration(delta))
				}
			}
		}
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("%s request body close failed: %s\n", logSectionTitle, err)
			//return
		}
	}(r.Body)

	RespondSuccess(w, "", responseData)
	return
}

// HandleGetAllWindows
// @Summary  Get all windows
// @Tags     windows
// @Accept      json
// @Produce     json
// @Success  200 {object} dto.Response{payload=repository.ObjectWindow}
// @Failure  400 {object} dto.Response
// @Router   /windows [get]
// @Security ApiKeyAuth
func HandleGetAllWindows(w http.ResponseWriter, r *http.Request) {
	logSectionTitle := loggerApiTitle + "[HandleGetAllWindows]"
	ctx := context.Background()

	windows, err := repository.GetAllWindows(ctx)

	if err != nil {
		log.Printf("%s %s\n", logSectionTitle, err)
		RespondError(w, "Get windows error!", windows, http.StatusBadRequest)
		return
	}

	RespondSuccess(w, "", windows)
	return

}
