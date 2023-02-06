package handler

import (
	"Go-Handson/entity"
	"Go-Handson/store"
	"Go-Handson/testutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListTask(t *testing.T) {
	type want struct {
		status  int
		rspFile string
	}
	tests := map[string]struct {
		tasks map[entity.TaskID]*entity.Task
		want  want
	}{
		"ok": {
			tasks: map[entity.TaskID]*entity.Task{
				1: {
					ID:     1,
					Title:  "test1",
					Status: "todo",
				},
				2: {
					ID:     2,
					Title:  "test2",
					Status: "done",
				},
			},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_task/ok_rsp.json.golden",
			},
		},
		"empty": {
			tasks: map[entity.TaskID]*entity.Task{},
			want: want{
				status:  http.StatusOK,
				rspFile: "testdata/list_task/empty_rsp.json.golden",
			},
		},
	}
	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/tasks", nil)

			sut := ListTask{Store: &store.TaskStore{Tasks: tt.tasks}}
			sut.ServeHTTP(w, r)

			resp := w.Result()
			testutil.AssertResponse(t,
				resp, tt.want.status, testutil.LoadFile(t, tt.want.rspFile),
			)
		})
	}
}
