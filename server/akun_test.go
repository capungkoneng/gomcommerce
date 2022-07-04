package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/capungkoneng/gomcommerce/db/mock"
	"github.com/stretchr/testify/require"

	db "github.com/capungkoneng/gomcommerce/db/sqlc"
	"github.com/capungkoneng/gomcommerce/util"
	"github.com/golang/mock/gomock"
)

func TestGetAkunApi(t *testing.T) {
	akun := randomAkun()

	testCases := []struct {
		name          string
		akunId        int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			akunId: akun.ID,
			buildStubs: func(store *mockdb.MockStore) {
				// Build stubs
				store.EXPECT().
					GetAuthor(gomock.Any(), gomock.Eq(akun.ID)).
					Times(1).
					Return(akun, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// check response
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAkun(t, recorder.Body, akun)
			},
		},
		{
			name:   "NotFound",
			akunId: akun.ID,
			buildStubs: func(store *mockdb.MockStore) {
				// Build stubs
				store.EXPECT().
					GetAuthor(gomock.Any(), gomock.Eq(akun.ID)).
					Times(1).
					Return(db.Akun{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// check response
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "InternalError",
			akunId: akun.ID,
			buildStubs: func(store *mockdb.MockStore) {
				// Build stubs
				store.EXPECT().
					GetAuthor(gomock.Any(), gomock.Eq(akun.ID)).
					Times(1).
					Return(db.Akun{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// check response
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:   "InvalidId",
			akunId: 0,
			buildStubs: func(store *mockdb.MockStore) {
				// Build stubs
				store.EXPECT().
					GetAuthor(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// check response
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/akun/%d", tc.akunId)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomAkun() db.Akun {
	return db.Akun{
		ID:       util.RandomInt(1, 100),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAkun(t *testing.T, body *bytes.Buffer, akun db.Akun) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAkun db.Akun
	err = json.Unmarshal(data, &gotAkun)
	require.Equal(t, akun, gotAkun)
}
